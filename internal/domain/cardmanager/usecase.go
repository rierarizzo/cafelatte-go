package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type DefaultManager struct {
	cardRepository CardRepository
}

func (manager DefaultManager) GetCardsByUserId(userID int) (
	[]domain.PaymentCard,
	*domain.AppError,
) {
	cards, appErr := manager.cardRepository.SelectCardsByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func (manager DefaultManager) AddUserCard(
	userId int,
	card domain.PaymentCard,
) (*domain.PaymentCard, *domain.AppError) {
	if appErr := validateCard(&card); appErr != nil {
		return nil, appErr
	}

	hash, appErr := crypt.HashText(card.Number)
	if appErr != nil {
		return nil, appErr
	}
	card.Number = hash

	hash, appErr = crypt.HashText(card.CVV)
	if appErr != nil {
		return nil, appErr
	}
	card.CVV = hash

	data, appErr := manager.cardRepository.InsertUserCard(userId, card)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return data, nil
}

func New(cardRepository CardRepository) *DefaultManager {
	return &DefaultManager{cardRepository}
}
