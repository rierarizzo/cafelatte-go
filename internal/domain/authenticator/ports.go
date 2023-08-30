package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

type Authenticator interface {
	SignUp(user domain.User) (*domain.AuthenticatedUser, *domain.AppError)
	SignIn(email, password string) (*domain.AuthenticatedUser, *domain.AppError)
}
