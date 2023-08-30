package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	addressRepository AddressRepository
}

func (m DefaultManager) GetAddressesByUserId(userId int) ([]domain.Address,
	*domain.AppError) {
	addresses, appErr := m.addressRepository.SelectAddressesByUserId(userId)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func (m DefaultManager) AddUserAddress(userId int,
	address domain.Address) (*domain.Address, *domain.AppError) {
	if appErr := validateAddress(&address); appErr != nil {
		return nil, appErr
	}

	data, appErr := m.addressRepository.InsertUserAddress(userId, address)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return data, nil
}

func New(addressRepository AddressRepository) *DefaultManager {
	return &DefaultManager{addressRepository}
}
