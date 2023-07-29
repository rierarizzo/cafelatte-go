package ports

import "github.com/rierarizzo/cafelatte/internal/domain/entities"

type IUserService interface {
	// SignUp es un método para registrar un usuario en el sistema.
	SignUp(user entities.User) (*entities.AuthorizedUser, error)
	// SignIn es un método para iniciar sesión en el sistema.
	SignIn(email, password string) (*entities.AuthorizedUser, error)
	// GetAllUsers es un método para obtener todos los usuarios registrados.
	GetAllUsers() ([]entities.User, error)
	// FindUserByID es un método para buscar un usuario por su ID.
	FindUserByID(userID int) (*entities.User, error)
	UpdateUser(userID int, user entities.User) error
}

type IUserRepository interface {
	SelectAllUsers() ([]entities.User, error)
	SelectUserByID(userID int) (*entities.User, error)
	SelectUserByEmail(email string) (*entities.User, error)
	InsertUser(user entities.User) (*entities.User, error)
	UpdateUser(userID int, user entities.User) error
}
