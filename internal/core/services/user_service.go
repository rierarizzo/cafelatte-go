package services

import (
	"github.com/rierarizzo/cafelatte/internal/core"
	"github.com/rierarizzo/cafelatte/internal/core/entities"
	"github.com/rierarizzo/cafelatte/internal/core/ports"
	"github.com/rierarizzo/cafelatte/internal/utils"
)

type UserService struct {
	userRepo ports.IUserRepository
}

func (us *UserService) SignUp(user entities.User) (*entities.AuthorizedUser, error) {
	hashedPassword, err := utils.HashText(user.Password)
	if err != nil {
		return nil, core.UnauthorizedUser
	}
	user.SetPassword(hashedPassword)

	retrievedUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		return nil, core.UnauthorizedUser
	}

	token, err := utils.CreateJWTToken(*retrievedUser)
	if err != nil {
		return nil, core.UnauthorizedUser
	}

	authorizedUser := entities.AuthorizedUser{
		UserInfo:    *retrievedUser,
		AccessToken: *token,
	}

	return &authorizedUser, nil
}

func (us *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	retrievedUser, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, core.UnauthorizedUser
	}

	if !utils.CheckTextHash(retrievedUser.Password, password) {
		return nil, core.UnauthorizedUser
	}

	authorizedUser := entities.AuthorizedUser{
		UserInfo:    *retrievedUser,
		AccessToken: "",
	}

	return &authorizedUser, nil
}

func (us *UserService) FindUserById(id int) (*entities.User, error) {
	return us.userRepo.GetUserById(id)
}

func NewUserService(userRepo ports.IUserRepository) *UserService {
	return &UserService{userRepo}
}
