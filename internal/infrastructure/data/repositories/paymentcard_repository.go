package repositories

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/mappers"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/data/models"
	"sync"
)

type PaymentCardRepository struct {
	db *sqlx.DB
}

var (
	selectCardError = errors.New("errors in selecting card(s)")
	insertCardError = errors.New("errors in inserting new card")
)

func (r PaymentCardRepository) SelectCardsByUserID(userID int) (
	[]entities.PaymentCard,
	error,
) {
	var cardsModel []models.PaymentCardModel

	query := "select * from userpaymentcard where UserID=? and Status=true"
	err := r.db.Select(&cardsModel, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, core.NewAppErrorWithType(core.NotFoundError)
		}
		return nil, core.NewAppError(
			errors.Join(selectCardError, err),
			core.RepositoryError,
		)
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
	returnRepoError := func(err error) error {
		return core.NewAppError(
			errors.Join(insertCardError, err),
			core.RepositoryError,
		)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return nil, returnRepoError(err)
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
		return nil, returnRepoError(err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, returnRepoError(err)
	}

	return cards, nil
}

func NewPaymentCardRepository(db *sqlx.DB) *PaymentCardRepository {
	return &PaymentCardRepository{db}
}
