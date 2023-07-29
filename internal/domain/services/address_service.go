package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
)

type AddressService struct {
	addressRepo ports.IAddressRepository
}

func (s AddressService) GetAddressesByUserID(userID int) (
	[]entities.Address,
	error,
) {
	addresses, err := s.addressRepo.SelectAddressesByUserID(userID)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != domain.NotFoundError) || !wrapped {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}

		return nil, err
	}

	return addresses, nil
}

func (s AddressService) AddUserAddresses(
	userID int,
	addresses []entities.Address,
) ([]entities.Address, error) {
	for _, v := range addresses {
		if err := v.ValidateAddress(); err != nil {
			return nil, domain.NewAppError(err, domain.ValidationError)
		}
	}

	addresses, err := s.addressRepo.InsertUserAddresses(userID, addresses)
	if err != nil {
		return nil, domain.NewAppError(err, domain.UnexpectedError)
	}

	return addresses, nil
}

func NewAddressService(addressRepo ports.IAddressRepository) *AddressService {
	return &AddressService{addressRepo}
}
