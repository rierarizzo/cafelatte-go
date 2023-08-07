package services

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/domain/validators"
)

type AddressService struct {
	addressRepo ports.IAddressRepository
}

func (s AddressService) GetAddressesByUserID(userID int) ([]entities.Address, *domain.AppError) {
	addresses, appErr := s.addressRepo.SelectAddressesByUserID(userID)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func (s AddressService) AddUserAddresses(userID int,
	addresses []entities.Address) ([]entities.Address, *domain.AppError) {
	for _, v := range addresses {
		if appErr := validators.ValidateAddress(&v); appErr != nil {
			return nil, appErr
		}
	}

	addresses, appErr := s.addressRepo.InsertUserAddresses(userID, addresses)
	if appErr != nil {
		if appErr.Type != domain.NotFoundError {
			return nil, domain.NewAppError(appErr, domain.UnexpectedError)
		}

		return nil, appErr
	}

	return addresses, nil
}

func NewAddressService(addressRepo ports.IAddressRepository) *AddressService {
	return &AddressService{addressRepo}
}
