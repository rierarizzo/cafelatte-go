package usermanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"mime/multipart"
)

type Manager interface {
	GetUsers() ([]domain.User, *domain.AppError)
	FindUserByID(userID int) (*domain.User, *domain.AppError)
	FindUserByEmail(email string) (*domain.User, *domain.AppError)
	UpdateUser(userID int, user domain.User) *domain.AppError
	DeleteUser(userID int) *domain.AppError
	UpdateProfilePic(userID int,
		pic *multipart.FileHeader) (string, *domain.AppError)
}

type UserRepository interface {
	SelectUsers() ([]domain.User, *domain.AppError)
	SelectUserByID(userID int) (*domain.User, *domain.AppError)
	SelectUserByEmail(email string) (*domain.User, *domain.AppError)
	InsertUser(user domain.User) (*domain.User, *domain.AppError)
	UpdateUser(userID int, user domain.User) *domain.AppError
	DeleteUser(userID int) *domain.AppError
}

type UserFilesRepository interface {
	UpdateProfilePic(userID int,
		pic *multipart.FileHeader, picname string) (string, *domain.AppError)
}
