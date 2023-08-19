package paymentcard

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
	"sync"
)

type Repository struct {
	db *sqlx.DB
}

func (r Repository) SelectCardsByUserID(userID int) ([]domain.PaymentCard, *domain.AppError) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var cardsModel []Model

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

	return fromModelsToCards(cardsModel), nil
}

func (r Repository) InsertUserPaymentCards(userID int,
	cards []domain.PaymentCard) ([]domain.PaymentCard, *domain.AppError) {
	rollbackAndError := func(tx *sqlx.Tx, err error) *domain.AppError {
		logrus.WithField(misc.RequestIDKey, request.ID()).Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertCardError, domain.RepositoryError)
	}

	tx, _ := r.db.Beginx()

	insertStmnt, err := tx.Prepare(`insert into UserPaymentCard (
                             Type, 
                             UserID, 
                             Company, 
                             HolderName, 
                             Number, 
                             ExpirationYear, 
                             ExpirationMonth, 
                             CVV) values (?,?,?,?,?,?,?,?)`)
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	sem := make(chan struct{}, 5)

	errCh := make(chan error, len(cards))
	var wg sync.WaitGroup

	for _, v := range cards {
		wg.Add(1)
		sem <- struct{}{}

		go func(card domain.PaymentCard) {
			defer func() {
				wg.Done()
				<-sem
			}()
			cardModel := fromCardToModel(card)

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
		return nil, rollbackAndError(tx, err)
	}

	_ = tx.Commit()

	return cards, nil
}

var (
	selectCardError = errors.New("select payment card error")
	insertCardError = errors.New("insert payment card error")
)

func NewPaymentCardRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}
