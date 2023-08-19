package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromUserToResponse(user domain.User) LoggedUserResponse {
	return LoggedUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.RoleCode,
	}
}

func fromAuthUserToResponse(authorizedUser domain.AuthorizedUser) AuthorizedUserResponse {
	loggedUserRes := fromUserToResponse(authorizedUser.User)

	return AuthorizedUserResponse{
		User:        loggedUserRes,
		AccessToken: authorizedUser.AccessToken,
	}
}

func fromSignUpRequestToUser(req SignUpRequest) domain.User {
	return domain.User{
		Username:    req.Username,
		Name:        req.Name,
		Surname:     req.Surname,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    req.Password,
		RoleCode:    req.RoleCode,
	}
}
