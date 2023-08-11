package paymentcard

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type Service struct {
	paymentCardRepo IPaymentCardRepository
}

func (s Service) GetCardsByUserID(userID int) ([]PaymentCard, *domain.AppError) {
	cards, appErr := s.paymentCardRepo.SelectCardsByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func (s Service) AddUserPaymentCard(userID int,
	cards []PaymentCard) ([]PaymentCard, *domain.AppError) {
	for k, v := range cards {
		if appErr := ValidatePaymentCard(&v); appErr != nil {
			return nil, appErr
		}

		if appErr := ValidateExpirationDate(&v); appErr != nil {
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

	cards, appErr := s.paymentCardRepo.InsertUserPaymentCards(userID, cards)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return cards, nil
}

func NewPaymentCardService(paymentCardRepo IPaymentCardRepository) *Service {
	return &Service{paymentCardRepo}
}
