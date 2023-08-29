package card

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/constants/misc"
	"github.com/rierarizzo/cafelatte/pkg/params/request"
	"github.com/sirupsen/logrus"
)

var (
	selectCardError = errors.New("select payment card error")
	insertCardError = errors.New("insert payment card error")
)

type Repository struct {
	db *sqlx.DB
}

func (repository Repository) SelectCardsByUserID(userId int) (
	[]domain.PaymentCard,
	*domain.AppError,
) {
	log := logrus.WithField(misc.RequestIDKey, request.ID())

	var cardsModel []Model

	query := "select * from UserPaymentCard where UserID=? and Status=true"
	err := repository.db.Select(&cardsModel, query, userId)
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

func (repository Repository) InsertUserCard(
	userId int,
	card domain.PaymentCard,
) (*domain.PaymentCard, *domain.AppError) {
	rollbackAndError := func(tx *sqlx.Tx, err error) *domain.AppError {
		logrus.WithField(misc.RequestIDKey, request.ID()).Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertCardError, domain.RepositoryError)
	}

	tx, _ := repository.db.Beginx()

	cardModel := fromCardToModel(card)
	result, err := tx.Exec(`insert into UserPaymentCard (
                             Type, 
                             UserID, 
                             Company, 
                             HolderName, 
                             Number, 
                             ExpirationYear, 
                             ExpirationMonth, 
                             CVV) values (?,?,?,?,?,?,?,?)`, cardModel.Type,
		userId,
		cardModel.Company, cardModel.HolderName, cardModel.Number,
		cardModel.ExpirationYear, cardModel.ExpirationMonth, cardModel.CVV)
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	cardID, _ := result.LastInsertId()
	card.ID = int(cardID)

	_ = tx.Commit()

	return &card, nil
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}
