package repositories

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

	query := "select * from UserPaymentCard where UserID=? and Status=true"
	err := r.db.Select(&cardsModel, query, userID)
	if err != nil {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			appErr := domain.NewAppErrorWithType(domain.NotFoundError)
			return nil, appErr
		}

		appErr := domain.NewAppError(selectCardError, domain.RepositoryError)
		return nil, appErr
	}

	return mappers.FromCardModelSliceToCardSlice(cardsModel), nil
}

func (r PaymentCardRepository) InsertUserPaymentCards(userID int,
	cards []entities.PaymentCard) ([]entities.PaymentCard, *domain.AppError) {
	log := logrus.WithField(constants.RequestIDKey, params.RequestID())

	returnError := func(err error) *domain.AppError {
		log.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertCardError, domain.RepositoryError)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, returnError(err)
	}

	insertStmnt, err := tx.Prepare(`insert into UserPaymentCard (
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
		return nil, returnError(err)
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
			cardModel := mappers.FromCardToCardModel(card)

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
		return nil, returnError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, returnError(err)
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
