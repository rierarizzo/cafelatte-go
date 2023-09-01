package card

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rierarizzo/cafelatte/internal/domain"
	sqlUtil "github.com/rierarizzo/cafelatte/pkg/utils/sql"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db}
}

func (r Repository) SelectCardsByUserId(userId int) ([]domain.PaymentCard, *domain.AppError) {
	var cardsModel []Model

	query := `
		SELECT * FROM UserPaymentCard WHERE UserId=? AND Status=TRUE
	`
	err := r.db.Select(&cardsModel, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.PaymentCard{}, nil
		}

		appErr := domain.NewAppError(err, domain.RepositoryError)
		return nil, appErr
	}

	cards := fromModelsToCards(cardsModel)
	return cards, nil
}

func (r Repository) InsertUserCard(userId int,
	card domain.PaymentCard) (*domain.PaymentCard, *domain.AppError) {
	tx, appErr := sqlUtil.StartTransaction(r.db)
	if appErr != nil {
		return nil, appErr
	}

	defer sqlUtil.RollbackIfPanic(tx)

	model := fromCardToModel(card)

	query := `
        INSERT INTO UserPaymentCard (Type, UserId, Company, HolderName, Number, 
        ExpirationYear, ExpirationMonth, CVV) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `
	result, appErr := sqlUtil.ExecWithTransaction(tx,
		query,
		model.Type,
		userId,
		model.Company,
		model.HolderName,
		model.Number,
		model.ExpirationYear,
		model.ExpirationMonth,
		model.CVV)
	if appErr != nil {
		return nil, appErr
	}

	cardId, appErr := sqlUtil.GetLastInsertedId(result)
	if appErr != nil {
		return nil, appErr
	}
	card.Id = cardId

	if appErr = sqlUtil.CommitTransaction(tx); appErr != nil {
		return nil, appErr
	}

	return &card, nil
}
