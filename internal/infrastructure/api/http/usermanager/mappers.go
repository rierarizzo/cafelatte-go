package usermanager

import (
	"github.com/rierarizzo/cafelatte/internal/domain"
	"strings"
)

func fromUserToResponse(user domain.User) Response {
	return Response{
		ID:           user.ID,
		CompleteName: strings.Join([]string{user.Name, user.Surname}, " "),
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.RoleCode,
	}
}

func fromUsersToResponse(users []domain.User) []Response {
	var res = make([]Response, 0)
	for _, v := range users {
		res = append(res, fromUserToResponse(v))
	}

	return res
}

func fromUpdateRequestToUser(updUserReq UpdateRequest) domain.User {
	return domain.User{
		Username:    updUserReq.Username,
		Name:        updUserReq.Name,
		Surname:     updUserReq.Surname,
		PhoneNumber: updUserReq.PhoneNumber,
	}
}
