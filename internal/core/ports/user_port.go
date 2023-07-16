package ports

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type IUserService interface {
	SignUp(user entities.User) (*entities.AuthorizedUser, error)
	SignIn(email, password string) (*entities.AuthorizedUser, error)
	FindUserById(id int) (*entities.User, error)
}

type IUserRepository interface {
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user entities.User) (*entities.User, error)
}
