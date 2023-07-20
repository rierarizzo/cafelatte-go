package ports

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type IUserService interface {
	// SignUp es un método para registrar un usuario en el sistema.
	SignUp(user entities.User) (*entities.AuthorizedUser, error)
	// SignIn es un método para iniciar sesión en el sistema.
	SignIn(email, password string) (*entities.AuthorizedUser, error)
	// GetAllUsers es un método para obtener todos los usuarios registrados.
	GetAllUsers() ([]entities.User, error)
	// FindUserById es un método para buscar un usuario por su ID.
	FindUserById(userID int) (*entities.User, error)
	UpdateUser(userID int, user entities.User) error
}

type IUserRepository interface {
	SelectAllUsers() ([]entities.User, error)
	SelectUserById(userID int) (*entities.User, error)
	SelectUserByEmail(email string) (*entities.User, error)
	InsertUser(user entities.User) (*entities.User, error)
	InsertUserPaymentCards(userID int, cards []entities.PaymentCard) ([]entities.PaymentCard, error)
	InsertUserAddresses(userID int, addresses []entities.Address) ([]entities.Address, error)
	UpdateUser(userID int, user entities.User) error
}
