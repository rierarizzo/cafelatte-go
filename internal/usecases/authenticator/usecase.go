package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/security/jsonwebtoken"
	"github.com/rierarizzo/cafelatte/internal/usecases/usermanager"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type DefaultAuthenticator struct {
	userRepository usermanager.UserRepository
}

func New(userRepository usermanager.UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepository}
}

func (a DefaultAuthenticator) SignUp(user domain.User) (*domain.AuthenticatedUser, *domain.AppError) {
	hashedPass, appErr := crypt.HashText(user.Password)
	if appErr != nil {
		return nil, appErr
	}
	user.SetPassword(hashedPass)

	returnedUser, appErr := a.userRepository.InsertUser(user)
	if appErr != nil {
		return nil, appErr
	}

	authorized, appErr := AuthorizeUser(*returnedUser)
	if appErr != nil {
		return nil, appErr
	}

	return authorized, nil
}

func (a DefaultAuthenticator) SignIn(email, password string) (*domain.AuthenticatedUser, *domain.AppError) {
	returnedUser, appErr := a.userRepository.SelectUserByEmail(email)
	if appErr != nil {
		if appErr.Type == domain.NotFoundError {
			return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
		}

		return nil, appErr
	}

	if !crypt.CheckTextHash(returnedUser.Password, password) {
		return nil, domain.NewAppErrorWithType(domain.NotAuthorizedError)
	}

	authorized, appErr := AuthorizeUser(*returnedUser)
	if appErr != nil {
		return nil, appErr
	}

	return authorized, nil
}

func AuthorizeUser(user domain.User) (*domain.AuthenticatedUser, *domain.AppError) {
	token, appErr := jsonwebtoken.CreateJWTToken(user)
	if appErr != nil {
		return nil, appErr
	}

	authorizedUser := domain.AuthenticatedUser{
		User:        user,
		AccessToken: token,
	}

	return &authorizedUser, nil
}
