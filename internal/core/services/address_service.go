package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	core "github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
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
		var coreErr *core.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != core.NotFoundError) || !wrapped {
			return nil, core.NewAppError(err, core.UnexpectedError)
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
			return nil, core.NewAppError(err, core.ValidationError)
		}
	}

	addresses, err := s.addressRepo.InsertUserAddresses(userID, addresses)
	if err != nil {
		return nil, core.NewAppError(err, core.UnexpectedError)
	}

	return addresses, nil
}

func NewAddressService(addressRepo ports.IAddressRepository) *AddressService {
	return &AddressService{addressRepo}
}
