package dto

import "github.com/rierarizzo/cafelatte/internal/core/entities"

type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (ur *UserResponse) LoadFromUserCore(user entities.User) {
	ur.ID = user.ID
	ur.Name = user.Name
	ur.Surname = user.Surname
}

type AuthResponse struct {
	User        UserResponse `json:"user"`
	AccessToken string       `json:"accesstoken"`
}

func (ar *AuthResponse) LoadFromAuthorizedUserCore(authorizedUser entities.AuthorizedUser) {
	var userResponse UserResponse
	userResponse.LoadFromUserCore(authorizedUser.User)
	ar.User = userResponse
	ar.AccessToken = authorizedUser.AccessToken
}
