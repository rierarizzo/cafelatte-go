package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	addressRepository AddressRepository
}

func (m DefaultManager) GetAddressesByUserID(userID int) ([]domain.Address, *domain.AppError) {
	addresses, appErr := m.addressRepository.SelectAddressesByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func (m DefaultManager) AddUserAddresses(userID int,
	addresses []domain.Address) ([]domain.Address, *domain.AppError) {
	for _, v := range addresses {
		if appErr := validateAddress(&v); appErr != nil {
			return nil, appErr
		}
	}

	addresses, appErr := m.addressRepository.InsertUserAddresses(userID, addresses)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func New(addressRepository AddressRepository) *DefaultManager {
	return &DefaultManager{addressRepository}
}
