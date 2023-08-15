package user

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"mime/multipart"
)

// IUserService represents an interface for a user service.
type IUserService interface {
	// GetUsers retrieves a list of users from the system and returns the list
	// of users if successful, along with any error encountered during the
	// process.
	GetUsers() ([]User, *domain.AppError)

	// FindUserByID retrieves a user from the system based on the provided user
	// ID and returns the user if found, along with any error encountered during
	// the process.
	FindUserByID(userID int) (*User, *domain.AppError)

	// FindUserByEmail retrieves a user from the system based on the
	// provided email and returns the user if found, along with any error
	// encountered during the process.
	FindUserByEmail(email string) (*User, *domain.AppError)

	CreateUser(user User) (*User, *domain.AppError)

	// UpdateUser updates the details of a user in the system based on the
	// provided user ID and user object and returns an error, if any,
	// encountered during the process.
	UpdateUser(userID int, user User) *domain.AppError

	DeleteUser(userID int) *domain.AppError
}

// IUserRepository represents an interface for a user repository.
type IUserRepository interface {
	// SelectUsers retrieves a list of users from the database and returns the
	// list of users if successful, along with any error encountered during the
	// process.
	SelectUsers() ([]User, *domain.AppError)

	// SelectUserByID retrieves a user from the database based on the provided
	// user ID and returns the user if found, along with any error encountered
	// during the process.
	SelectUserByID(userID int) (*User, *domain.AppError)

	// SelectUserByEmail retrieves a user from the database based on the
	// provided email and returns the user if found, along with any error
	// encountered during the process.
	SelectUserByEmail(email string) (*User, *domain.AppError)

	// InsertUser inserts a new user into the database and returns the inserted
	// user if successful, along with any error encountered during the process.
	InsertUser(user User) (*User, *domain.AppError)

	// UpdateUser updates the details of a user in the database based on the
	// provided user ID and user object and returns an error, if any,
	// encountered during the process.
	UpdateUser(userID int, user User) *domain.AppError

	DeleteUser(userID int) *domain.AppError
}

type IUserFilesRepository interface {
	UpdateProfilePic(userID int,
		pic *multipart.FileHeader) (string, *domain.AppError)
}
