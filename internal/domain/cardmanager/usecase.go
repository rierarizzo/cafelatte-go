package cardmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type DefaultManager struct {
	cardRepository CardRepository
}

func (m DefaultManager) GetCardsByUserID(userID int) ([]domain.PaymentCard, *domain.AppError) {
	cards, appErr := m.cardRepository.SelectCardsByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func (m DefaultManager) AddUserPaymentCard(userID int,
	cards []domain.PaymentCard) ([]domain.PaymentCard, *domain.AppError) {
	for k, v := range cards {
		if appErr := validateCard(&v); appErr != nil {
			return nil, appErr
		}

		hash, appErr := crypt.HashText(v.Number)
		if appErr != nil {
			return nil, appErr
		}
		cards[k].Number = hash

		hash, appErr = crypt.HashText(v.CVV)
		if appErr != nil {
			return nil, appErr
		}
		cards[k].CVV = hash
	}

	cards, appErr := m.cardRepository.InsertUserPaymentCards(userID, cards)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func NewDefaultManager(cardRepository CardRepository) *DefaultManager {
	return &DefaultManager{cardRepository}
}
