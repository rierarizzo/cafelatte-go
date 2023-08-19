package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Authenticator interface {
	SignUp(user domain.User) (*domain.AuthorizedUser, *domain.AppError)
	SignIn(email, password string) (*domain.AuthorizedUser, *domain.AppError)
}
