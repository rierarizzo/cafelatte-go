package authenticator

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
)

func fromAuthUserToResponse(authorizedUser domain.AuthenticatedUser) AuthenticatedResponse {
	return AuthenticatedResponse{
		User: struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}{
			Id:       authorizedUser.User.Id,
			Username: authorizedUser.User.Username,
			Email:    authorizedUser.User.Email,
			Role:     authorizedUser.User.RoleCode,
		},
		AccessToken: authorizedUser.AccessToken,
	}
}

func fromRequestToUser(req UserSignup) domain.User {
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
