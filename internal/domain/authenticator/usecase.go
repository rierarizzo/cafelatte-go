package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"github.com/rierarizzo/cafelatte/internal/domain/usermanager"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/security/jsonwebtoken"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type DefaultAuthenticator struct {
	userRepository usermanager.UserRepository
}

func (a DefaultAuthenticator) SignUp(user domain.User) (*domain.AuthorizedUser, *domain.AppError) {
	if appErr := validateUser(&user); appErr != nil {
		return nil, appErr
	}

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

func (a DefaultAuthenticator) SignIn(email, password string) (*domain.AuthorizedUser, *domain.AppError) {
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

func AuthorizeUser(user domain.User) (*domain.AuthorizedUser, *domain.AppError) {
	token, appErr := jsonwebtoken.CreateJWTToken(user)
	if appErr != nil {
		return nil, appErr
	}

	authorizedUser := domain.AuthorizedUser{
		User:        user,
		AccessToken: token,
	}

	return &authorizedUser, nil
}

func New(userRepository usermanager.UserRepository) *DefaultAuthenticator {
	return &DefaultAuthenticator{userRepository}
}
