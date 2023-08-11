package authenticate

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	domainUser "github.com/rierarizzo/cafelatte/internal/domain/user"
	"github.com/rierarizzo/cafelatte/internal/infrastructure/security/jsonwebtoken"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type Usecase struct {
	userService domainUser.IUserService
}

// SignUp registers a new user in the system and returns an AuthorizedUser
// along with any error encountered during the process.
func (a Usecase) SignUp(user domainUser.User) (*AuthorizedUser, *domain.AppError) {
	if appErr := domainUser.ValidateUser(&user); appErr != nil {
		return nil, appErr
	}

	if appErr := user.HashPassword(); appErr != nil {
		return nil, appErr
	}

	returnedUser, appErr := a.userService.CreateUser(user)
	if appErr != nil {
		return nil, appErr
	}

	authorized, appErr := AuthorizeUser(*returnedUser)
	if appErr != nil {
		return nil, appErr
	}

	return authorized, nil
}

func (a Usecase) SignIn(email string,
	password string) (*AuthorizedUser, *domain.AppError) {
	returnedUser, appErr := a.userService.FindUserByEmail(email)
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

func AuthorizeUser(user domainUser.User) (*AuthorizedUser, *domain.AppError) {
	token, appErr := jsonwebtoken.CreateJWTToken(user)
	if appErr != nil {
		return nil, appErr
	}

	authorizedUser := AuthorizedUser{
		User:        user,
		AccessToken: token,
	}

	return &authorizedUser, nil
}

func NewAuthenticateUsecase(userService domainUser.IUserService) *Usecase {
	return &Usecase{userService}
}
