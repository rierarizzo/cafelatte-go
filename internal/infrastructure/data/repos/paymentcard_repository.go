package repos

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/constants"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"github.com/rierarizzo/cafelatte/internal/params"
	"github.com/sirupsen/logrus"
	"sync"
)

type PaymentCardRepository struct {
	db *sqlx.DB
}

func (r PaymentCardRepository) SelectCardsByUserID(userID int) ([]entities.PaymentCard, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	var cardsModel []models.PaymentCardModel

	query := "select * from userpaymentcard where UserID=? and Status=true"
	err := r.db.Select(&cardsModel, query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return nil, domain.NewAppError(selectCardError, domain.RepositoryError)
	}

	var cards []entities.PaymentCard
	for _, v := range cardsModel {
		cards = append(cards, mappers.FromPaymentCardModelToPaymentCard(v))
	}

	return cards, nil
}

func (r PaymentCardRepository) InsertUserPaymentCards(userID int,
	cards []entities.PaymentCard) ([]entities.PaymentCard, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	returnRepoError := func(err error) *domain.AppError {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertCardError, domain.RepositoryError)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, returnRepoError(err)
	}

	insertStmnt, err := tx.Prepare(`insert into userpaymentcard (
                             Type, 
                             UserID, 
                             Company, 
                             HolderName, 
                             Number, 
                             ExpirationYear, 
                             ExpirationMonth, 
                             CVV
            ) values (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, returnRepoError(err)
	}

	sem := make(chan struct{}, 5)

	errCh := make(chan error, len(cards))
	var wg sync.WaitGroup

	for _, v := range cards {
		wg.Add(1)
		sem <- struct{}{}

		go func(card entities.PaymentCard) {
			defer func() {
				wg.Done()
				<-sem
			}()
			cardModel := mappers.FromPaymentCardToPaymentCardModel(card)

			result, err := insertStmnt.Exec(cardModel.Type, userID,
				cardModel.Company, cardModel.HolderName, cardModel.Number,
				cardModel.ExpirationYear, cardModel.ExpirationMonth,
				cardModel.CVV)
			if err != nil {
				errCh <- err
				return
			}

			cardID, _ := result.LastInsertId()
			card.ID = int(cardID)
		}(v)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		_ = tx.Rollback()
		return nil, returnRepoError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, returnRepoError(err)
	}

	return cards, nil
}

var (
	selectCardError = errors.New("select payment card error")
	insertCardError = errors.New("insert payment card error")
)

func NewPaymentCardRepository(db *sqlx.DB) *PaymentCardRepository {
	return &PaymentCardRepository{db}
}
