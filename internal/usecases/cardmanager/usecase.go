package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type DefaultManager struct {
	cardRepository CardRepository
}

func New(cardRepository CardRepository) *DefaultManager {
	return &DefaultManager{cardRepository}
}

func (m DefaultManager) GetCardsByUserId(userId int) ([]domain.PaymentCard,
	*domain.AppError) {
	cards, appErr := m.cardRepository.SelectCardsByUserId(userId)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func (m DefaultManager) AddUserCard(userId int,
	card domain.PaymentCard) (*domain.PaymentCard, *domain.AppError) {
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

	data, appErr := m.cardRepository.InsertUserCard(userId, card)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return data, nil
}
