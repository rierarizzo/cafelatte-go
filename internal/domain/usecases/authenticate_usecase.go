package usecases

import (
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/security"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type AuthenticateUsecase struct {
	userService ports.IUserService
}

// SignUp registers a new user in the system and returns an AuthorizedUser
// along with any error encountered during the process.
func (a AuthenticateUsecase) SignUp(user entities.User) (*entities.AuthorizedUser, *domain.AppError) {
	if appErr := user.ValidateUser(); appErr != nil {
		return nil, appErr
	}

	if appErr := user.HashPassword(); appErr != nil {
		return nil, appErr
	}

	rUser, appErr := a.userService.CreateUser(user)
	if appErr != nil {
		return nil, appErr
	}

	authorized, appErr := AuthorizeUser(*rUser)
	if appErr != nil {
		return nil, appErr
	}

	return authorized, nil
}

func (a AuthenticateUsecase) SignIn(email string,
	password string) (*entities.AuthorizedUser, *domain.AppError) {
	rUser, appErr := a.userService.FindUserByEmail(email)
	if appErr != nil {
		if appErr.Type == domain.NotFoundError {
			return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
		}

		return nil, appErr
	}

	if !utils.CheckTextHash(rUser.Password, password) {
		return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
	}

	authorized, appErr := AuthorizeUser(*rUser)
	if appErr != nil {
		return nil, appErr
	}

	return authorized, nil
}

func AuthorizeUser(user entities.User) (*entities.AuthorizedUser, *domain.AppError) {
	token, appErr := security.CreateJWTToken(user)
	if appErr != nil {
		return nil, appErr
	}

	authorizedUser := entities.AuthorizedUser{
		User:        user,
		AccessToken: token,
	}

	return &authorizedUser, nil
}

func NewAuthenticateUsecase(userService ports.IUserService) *AuthenticateUsecase {
	return &AuthenticateUsecase{userService}
}
