package addressmanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetAddressesByUserId(userId int) ([]domain.Address, *domain.AppError)
	AddUserAddress(userId int, address domain.Address) (*domain.Address,
		*domain.AppError)
}

type AddressRepository interface {
	SelectAddressesByUserId(userId int) ([]domain.Address, *domain.AppError)
	InsertUserAddress(userId int, address domain.Address) (*domain.Address,
		*domain.AppError)
	SelectCityNameById(id int) (string, *domain.AppError)
	SelectProvinceNameById(id int) (string, *domain.AppError)
}
