package authenticate

import (
	"github.com/rierarizzo/cafelatte/internal/domain/authenticate"
	"github.com/rierarizzo/cafelatte/internal/domain/user"
)

func FromUserToLoggedUser(user user.User) LoggedUserResponse {
	return LoggedUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.RoleCode,
	}
}

func FromAuthorizedUserToAuthorizationRes(authorizedUser authenticate.AuthorizedUser) AuthorizedUserResponse {
	loggedUserRes := FromUserToLoggedUser(authorizedUser.User)

	return AuthorizedUserResponse{
		User:        loggedUserRes,
		AccessToken: authorizedUser.AccessToken,
	}
}
