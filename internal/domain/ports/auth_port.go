package ports

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
)

type IAuthService interface {
	// SignUp registers a new user in the system and returns an AuthorizedUser
	// along with any error encountered during the process.
	SignUp(user entities.User) (*entities.AuthorizedUser, error)

	// SignIn authenticates a user with the provided email and password and
	// returns an AuthorizedUser if the authentication is successful, along
	// with any error encountered during the process.
	SignIn(email, password string) (*entities.AuthorizedUser, error)
}
