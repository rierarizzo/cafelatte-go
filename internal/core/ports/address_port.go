package ports

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type IAddressService interface {
	GetAddressByID(userID int, addressID int) (*entities.Address, error)
	GetAddressesByUserID(userID int) ([]entities.Address, error)
	AddUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error)
}

type IAddressRepository interface {
	SelectAddressByID(userID int, addressID int) (*entities.Address, error)
	SelectAddressesByUserID(userID int) ([]entities.Address, error)
	SelectCityNameByCityID(cityID int) (string, error)
	SelectProvinceNameByProvinceID(cityID int) (string, error)
	InsertUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error)
}
