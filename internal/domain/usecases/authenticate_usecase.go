package usecases

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/infra/security"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type AuthenticateUsecase struct {
	userService ports.IUserService
}

// SignUp registers a new user in the system and returns an AuthorizedUser
// along with any error encountered during the process.
func (a AuthenticateUsecase) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	// Validating user
	if err := user.ValidateUser(); err != nil {
		return nil, domain.NewAppError(err, domain.ValidationError)
	}
	// Hashing password
	if err := user.HashPassword(); err != nil {
		return nil, domain.NewAppError(err, domain.HashGenerationError)
	}
	// Inserting user to database
	retrUser, err := a.userService.CreateUser(user)
	if err != nil {
		return nil, domain.NewAppError(err, domain.UnexpectedError)
	}
	// Generating JWT
	authUser, err := AuthorizeUser(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	return authUser, nil
}

func (a AuthenticateUsecase) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	// Select user from database
	retrUser, err := a.userService.FindUserByEmail(email)
	if err != nil {
		var coreErr *domain.AppError
		converted := errors.As(err, &coreErr)
		if converted && coreErr.Type == domain.NotFoundError {
			return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
		} else {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}
	}
	// Verifying password hash
	if !utils.CheckTextHash(retrUser.Password, password) {
		return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
	}
	// Creating JWT
	authUser, err := AuthorizeUser(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	return authUser, nil
}

func AuthorizeUser(user entities.User) (*entities.AuthorizedUser, error) {
	token, err := security.CreateJWTToken(user)
	if err != nil {
		return nil, err
	}

	authorizedUser := entities.AuthorizedUser{
		User:        user,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func NewAuthenticateUsecase(userService ports.IUserService) *AuthenticateUsecase {
	return &AuthenticateUsecase{userService}
}
