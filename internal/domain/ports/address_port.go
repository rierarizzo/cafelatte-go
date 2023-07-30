package ports

import "github.com/rierarizzo/cafelatte/internal/domain/entities"

type IAddressService interface {
	GetAddressesByUserID(userID int) ([]entities.Address, error)
	AddUserAddresses(userID int,
		addresses []entities.Address) ([]entities.Address, error)
}

type IAddressRepository interface {
	SelectAddressesByUserID(userID int) ([]entities.Address, error)
	InsertUserAddresses(userID int,
		addresses []entities.Address) ([]entities.Address, error)
	SelectCityNameByCityID(cityID int) (string, error)
	SelectProvinceNameByProvinceID(cityID int) (string, error)
}
