package authenticate

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/user"
)

type IAuthenticateUsecase interface {
	// SignUp registers a new user in the system and returns an AuthorizedUser
	// along with any error encountered during the process.
	SignUp(user user.User) (*AuthorizedUser, *domain.AppError)

	// SignIn authenticates a user with the provided email and password and
	// returns an AuthorizedUser if the authentication is successful, along
	// with any error encountered during the process.
	SignIn(email string,
		password string) (*AuthorizedUser, *domain.AppError)
}
