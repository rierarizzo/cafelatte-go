package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	coreErrors "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"sync"
)

type PaymentCardRepository struct {
	db *sqlx.DB
}

func (r PaymentCardRepository) SelectCardsByUserID(userID int) (
	[]entities.PaymentCard,
	error,
) {
	var cardsModel []models.PaymentCardModel

	query := "select * from userpaymentcard where UserID=? and Status=true"
	err := r.db.Select(&cardsModel, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, coreErrors.WrapError(
				coreErrors.ErrRecordNotFound,
				err.Error(),
			)
		}
		return nil, coreErrors.WrapError(coreErrors.ErrUnexpected, err.Error())
	}

	var cards []entities.PaymentCard
	for _, v := range cardsModel {
		cards = append(cards, *mappers.FromPaymentCardModelToPaymentCard(v))
	}

	return cards, nil
}

func (r PaymentCardRepository) InsertUserPaymentCards(
	userID int,
	cards []entities.PaymentCard,
) ([]entities.PaymentCard, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, coreErrors.WrapError(coreErrors.ErrUnexpected, err.Error())
	}

	insertStmnt, err := tx.Prepare(
		`insert into userpaymentcard (
                             Type, 
                             UserID, 
                             Company, 
                             HolderName, 
                             Number, 
                             ExpirationYear, 
                             ExpirationMonth, 
                             CVV
            ) values (?,?,?,?,?,?,?,?)`,
	)
	if err != nil {
		return nil, coreErrors.WrapError(coreErrors.ErrUnexpected, err.Error())
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

			result, err := insertStmnt.Exec(
				cardModel.Type,
				userID,
				cardModel.Company,
				cardModel.HolderName,
				cardModel.Number,
				cardModel.ExpirationYear,
				cardModel.ExpirationMonth,
				cardModel.CVV,
			)
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
		return nil, coreErrors.WrapError(coreErrors.ErrUnexpected, err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, coreErrors.WrapError(coreErrors.ErrUnexpected, err.Error())
	}

	return cards, nil
}

func NewPaymentCardRepository(db *sqlx.DB) *PaymentCardRepository {
	return &PaymentCardRepository{db}
}
