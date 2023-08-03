package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IAddressService interface {
	GetAddressesByUserID(userID int) ([]entities.Address, *domain.AppError)
	AddUserAddresses(userID int,
		addresses []entities.Address) ([]entities.Address, *domain.AppError)
}

type IAddressRepository interface {
	SelectAddressesByUserID(userID int) ([]entities.Address, *domain.AppError)
	InsertUserAddresses(userID int,
		addresses []entities.Address) ([]entities.Address, *domain.AppError)
	SelectCityNameByCityID(cityID int) (string, *domain.AppError)
	SelectProvinceNameByProvinceID(cityID int) (string, *domain.AppError)
}
