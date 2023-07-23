package services

import (
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
)

type AddressService struct {
	addressRepo ports.IAddressRepository
}

func (a AddressService) GetAddressByID(userID int, addressID int) (*entities.Address, error) {
	return a.addressRepo.SelectAddressByID(userID, addressID)
}

func (a AddressService) GetAddressesByUserID(userID int) ([]entities.Address, error) {
	return a.addressRepo.SelectAddressesByUserID(userID)
}

func (a AddressService) AddUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error) {
	for _, v := range addresses {
		if err := v.ValidateAddress(); err != nil {
			return nil, err
		}
	}

	return a.addressRepo.InsertUserAddresses(userID, addresses)
}

func NewAddressService(addressRepo ports.IAddressRepository) *AddressService {
	return &AddressService{addressRepo}
}
