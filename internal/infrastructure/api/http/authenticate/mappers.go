package authenticate

import (
	"github.com/rierarizzo/cafelatte/internal/domain/authenticate"
	"github.com/rierarizzo/cafelatte/internal/domain/user"
)

func fromUserToResponse(user user.User) LoggedUserResponse {
	return LoggedUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.RoleCode,
	}
}

func fromAuthUserToResponse(authorizedUser authenticate.AuthorizedUser) AuthorizedUserResponse {
	loggedUserRes := fromUserToResponse(authorizedUser.User)

	return AuthorizedUserResponse{
		User:        loggedUserRes,
		AccessToken: authorizedUser.AccessToken,
	}
}

func fromSignUpRequestToUser(req SignUpRequest) user.User {
	return user.User{
		Username:    req.Username,
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		RoleCode:    req.RoleCode,
	}
}
