package ports

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type IUserService interface {
	// SignUp es un método para registrar un usuario en la base de datos.
	SignUp(user entities.User) (*entities.AuthorizedUser, error)
	// SignIn es un método para iniciar sesión.
	SignIn(email, password string) (*entities.AuthorizedUser, error)
	// GetAllUsers es un método para obtener todos los usuarios registrados
	// en la base de datos.
	GetAllUsers() ([]entities.User, error)
	// FindUserById es un método para buscar un usuario por su ID.
	FindUserById(id int) (*entities.User, error)
}

type IUserRepository interface {
	GetAllUsers() ([]entities.User, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user entities.User) (*entities.User, error)
}
