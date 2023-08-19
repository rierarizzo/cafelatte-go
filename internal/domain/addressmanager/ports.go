package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetAddressesByUserID(userID int) ([]domain.Address, *domain.AppError)
	AddUserAddresses(userID int,
		addresses []domain.Address) ([]domain.Address, *domain.AppError)
}

type AddressRepository interface {
	SelectAddressesByUserID(userID int) ([]domain.Address, *domain.AppError)
	InsertUserAddresses(userID int,
		addresses []domain.Address) ([]domain.Address, *domain.AppError)
	SelectCityNameByCityID(cityID int) (string, *domain.AppError)
	SelectProvinceNameByProvinceID(cityID int) (string, *domain.AppError)
}
