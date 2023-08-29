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
	selectCardError = errors.New("select card error")
	insertCardError = errors.New("insert card error")
)

type Repository struct {
	db *sqlx.DB
}

func (repository Repository) SelectCardsByUserId(userId int) (
	[]domain.PaymentCard,
	*domain.AppError,
) {
	log := logrus.WithField(misc.RequestIdKey, request.Id())

	var cardsModel []Model

	query := "select * from UserPaymentCard where UserId=? and Status=true"
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
		logrus.WithField(misc.RequestIdKey, request.Id()).Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return domain.NewAppErrorWithType(domain.NotFoundError)
		}

		return domain.NewAppError(insertCardError, domain.RepositoryError)
	}

	tx, _ := repository.db.Beginx()

	cardModel := fromCardToModel(card)
	query := `insert into UserPaymentCard (Type, UserId, Company, HolderName, Number, 
        ExpirationYear, ExpirationMonth, CVV) values (?,?,?,?,?,?,?,?)`

	result, err := tx.Exec(query, cardModel.Type, userId, cardModel.Company,
		cardModel.HolderName, cardModel.Number, cardModel.ExpirationYear,
		cardModel.ExpirationMonth, cardModel.CVV)
	if err != nil {
		return nil, rollbackAndError(tx, err)
	}

	cardId, _ := result.LastInsertId()
	card.Id = int(cardId)

	_ = tx.Commit()

	return &card, nil
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}
