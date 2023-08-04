package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
)

// IUserService represents an interface for a user service.
type IUserService interface {
	// GetUsers retrieves a list of users from the system and returns the list
	// of users if successful, along with any error encountered during the
	// process.
	GetUsers() ([]entities.User, *domain.AppError)

	// FindUserByID retrieves a user from the system based on the provided user
	// ID and returns the user if found, along with any error encountered during
	// the process.
	FindUserByID(userID int) (*entities.User, *domain.AppError)

	// FindUserByEmail retrieves a user from the system based on the
	// provided email and returns the user if found, along with any error
	// encountered during the process.
	FindUserByEmail(email string) (*entities.User, *domain.AppError)

	CreateUser(user entities.User) (*entities.User, *domain.AppError)

	// UpdateUser updates the details of a user in the system based on the
	// provided user ID and user object and returns an error, if any,
	// encountered during the process.
	UpdateUser(userID int, user entities.User) *domain.AppError
}

// IUserRepository represents an interface for a user repository.
type IUserRepository interface {
	// SelectUsers retrieves a list of users from the database and returns the
	// list of users if successful, along with any error encountered during the
	// process.
	SelectUsers() ([]entities.User, *domain.AppError)

	// SelectUserByID retrieves a user from the database based on the provided
	// user ID and returns the user if found, along with any error encountered
	// during the process.
	SelectUserByID(userID int) (*entities.User, *domain.AppError)

	// SelectUserByEmail retrieves a user from the database based on the
	// provided email and returns the user if found, along with any error
	// encountered during the process.
	SelectUserByEmail(email string) (*entities.User, *domain.AppError)

	// InsertUser inserts a new user into the database and returns the inserted
	// user if successful, along with any error encountered during the process.
	InsertUser(user entities.User) (*entities.User, *domain.AppError)

	// UpdateUser updates the details of a user in the database based on the
	// provided user ID and user object and returns an error, if any,
	// encountered during the process.
	UpdateUser(userID int, user entities.User) *domain.AppError
}
