package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type DefaultManager struct {
	addressRepository AddressRepository
}

func (manager DefaultManager) GetAddressesByUserId(userId int) (
	[]domain.Address,
	*domain.AppError,
) {
	addresses, appErr := manager.addressRepository.SelectAddressesByUserId(userId)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func (manager DefaultManager) AddUserAddress(
	userId int,
	address domain.Address,
) (*domain.Address, *domain.AppError) {
	if appErr := validateAddress(&address); appErr != nil {
		return nil, appErr
	}

	data, appErr := manager.addressRepository.InsertUserAddress(userId, address)
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
