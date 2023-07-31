package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type AuthService struct {
	userRepo ports.IUserRepo
}

// SignUp registers a new user in the system and returns an AuthorizedUser
// along with any error encountered during the process.
func (s *AuthService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	// Validating user
	if err := user.ValidateUser(); err != nil {
		return nil, domain.NewAppError(err, domain.ValidationError)
	}
	// Hashing password
	if err := user.HashPassword(); err != nil {
		return nil, domain.NewAppError(err, domain.HashGenerationError)
	}
	// Inserting user to database
	retrUser, err := s.userRepo.InsertUser(user)
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

// SignIn authenticates a user with the provided email and password and
// returns an AuthorizedUser if the authentication is successful, along
// with any error encountered during the process.
func (s *AuthService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	// Select user from database
	retrUser, err := s.userRepo.SelectUserByEmail(email)
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

func NewAuthService(userRepo ports.IUserRepo) *AuthService {
	return &AuthService{userRepo}
}
