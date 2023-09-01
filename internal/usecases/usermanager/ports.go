package usermanager

import (
	"mime/multipart"

	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Manager interface {
	GetUsers() ([]domain.User, *domain.AppError)
	FindUserById(id int) (*domain.User, *domain.AppError)
	FindUserByEmail(email string) (*domain.User, *domain.AppError)
	UpdateUserById(id int, user domain.User) *domain.AppError
	DeleteUserById(id int) *domain.AppError
	UpdateProfilePicById(id int, pic *multipart.FileHeader) (string,
		*domain.AppError)
}

type UserRepository interface {
	SelectUsers() ([]domain.User, *domain.AppError)
	SelectUserById(id int) (*domain.User, *domain.AppError)
	SelectUserByEmail(email string) (*domain.User, *domain.AppError)
	InsertUser(user domain.User) (*domain.User, *domain.AppError)
	UpdateUserById(id int, user domain.User) *domain.AppError
	DeleteUserById(id int) *domain.AppError
}

type UserFilesRepository interface {
	UpdateProfilePicById(id int, pic *multipart.FileHeader,
		picname string) (string, *domain.AppError)
}
