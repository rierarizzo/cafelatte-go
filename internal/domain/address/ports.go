package address

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

type IAddressService interface {
	GetAddressesByUserID(userID int) ([]Address, *domain.AppError)
	AddUserAddresses(userID int,
		addresses []Address) ([]Address, *domain.AppError)
}

type IAddressRepository interface {
	SelectAddressesByUserID(userID int) ([]Address, *domain.AppError)
	InsertUserAddresses(userID int,
		addresses []Address) ([]Address, *domain.AppError)
	SelectCityNameByCityID(cityID int) (string, *domain.AppError)
	SelectProvinceNameByProvinceID(cityID int) (string, *domain.AppError)
}
