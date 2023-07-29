package services

import (
	"errors"
	"github.com/rierarizzo/cafelatte/internal/domain/entities"
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/internal/domain/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

// UserService represents a user service implementation.
type UserService struct {
	userRepo ports.IUserRepo
}

// SignUp registers a new user in the system and returns an AuthorizedUser
// along with any error encountered during the process.
func (s *UserService) SignUp(user entities.User) (
	*entities.AuthorizedUser,
	error,
) {
	// Validating user
	if err := user.ValidateUser(); err != nil {
		return nil, domain.NewAppError(err, domain.ValidationError)
	}

	// Hashing password
	hash, err := utils.HashText(user.Password)
	if err != nil {
		return nil, domain.NewAppErrorWithType(domain.HashGenerationError)
	}
	user.Password = hash

	// Inserting user to database
	retrUser, err := s.userRepo.InsertUser(user)
	if err != nil {
		return nil, domain.NewAppError(err, domain.UnexpectedError)
	}

	// Generating JWT
	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	authorizedUser := entities.NewAuthorizedUser(*retrUser, *token)

	return authorizedUser, nil
}

// SignIn authenticates a user with the provided email and password and
// returns an AuthorizedUser if the authentication is successful, along
// with any error encountered during the process.
func (s *UserService) SignIn(email, password string) (
	*entities.AuthorizedUser,
	error,
) {
	// Select user from database
	retrUser, err := s.userRepo.SelectUserByEmail(email)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == domain.NotFoundError {
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
	token, err := utils.CreateJWTToken(*retrUser)
	if err != nil {
		return nil, domain.NewAppError(err, domain.TokenGenerationError)
	}

	return entities.NewAuthorizedUser(*retrUser, *token), nil
}

// GetUsers retrieves a list of users from the system and returns the list
// of users if successful, along with any error encountered during the
// process.
func (s *UserService) GetUsers() ([]entities.User, error) {
	users, err := s.userRepo.SelectUsers()
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if wrapped && coreErr.Type == domain.NotFoundError {
			return []entities.User{}, nil
		} else {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}
	}

	return users, nil
}

// FindUserByID retrieves a user from the system based on the provided user
// ID and returns the user if found, along with any error encountered during
// the process.
func (s *UserService) FindUserByID(id int) (*entities.User, error) {
	user, err := s.userRepo.SelectUserByID(id)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != domain.NotFoundError) || !wrapped {
			return nil, domain.NewAppError(err, domain.UnexpectedError)
		}

		return nil, err
	}

	return user, nil
}

// UpdateUser updates the details of a user in the system based on the
// provided user ID and user object and returns an error, if any,
// encountered during the process.
func (s *UserService) UpdateUser(userID int, user entities.User) error {
	err := s.userRepo.UpdateUser(userID, user)
	if err != nil {
		var coreErr *domain.AppError
		wrapped := errors.As(err, &coreErr)
		if (wrapped && coreErr.Type != domain.NotFoundError) || !wrapped {
			return domain.NewAppError(err, domain.UnexpectedError)
		}

		return err
	}

	return nil
}

func NewUserService(userRepo ports.IUserRepo) *UserService {
	return &UserService{userRepo}
}