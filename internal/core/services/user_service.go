package services

import (
	"errors"
	"log"

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
		log.Println(err)
		return nil, err
	}

	user.Password = hashedPassword

	retrievedUser, err := us.userRepo.CreateUser(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	authorizedUser := entities.AuthorizedUser{
		UserInfo:    *retrievedUser,
		AccessToken: "",
	}

	return &authorizedUser, nil
}

func (us *UserService) SignIn(email, password string) (*entities.AuthorizedUser, error) {
	retrievedUser, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unauthorized")
	}

	if !utils.CheckTextHash(retrievedUser.Password, password) {
		log.Println("Unauthorized user")
		return nil, errors.New("unauthorized")
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
