package services

import (
	"fmt"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/errors"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
)

type AddressService struct {
	addressRepo ports.IAddressRepository
}

func (s AddressService) GetAddressesByUserID(userID int) ([]entities.Address, error) {
	return s.addressRepo.SelectAddressesByUserID(userID)
}

func (s AddressService) AddUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error) {
	for _, v := range addresses {
		if err := v.ValidateAddress(); err != nil {
			return nil, errors.WrapError(err,
				fmt.Sprintf("address with detail '%s' is not valid", v.Detail))
		}
	}

	return s.addressRepo.InsertUserAddresses(userID, addresses)
}

func NewAddressService(addressRepo ports.IAddressRepository) *AddressService {
	return &AddressService{addressRepo}
}
