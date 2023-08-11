package user

import (
	"github.com/rierarizzo/cafelatte/internal/domain/user"
	"strings"
)

func fromUserToResponse(user user.User) Response {
	return Response{
		ID:           user.ID,
		CompleteName: strings.Join([]string{user.Name, user.Surname}, " "),
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.RoleCode,
	}
}

func fromUsersToResponse(users []user.User) []Response {
	var res = make([]Response, 0)
	for _, v := range users {
		res = append(res, fromUserToResponse(v))
	}

	return res
}

func fromUpdateRequestToUser(updUserReq UpdateRequest) user.User {
	return user.User{
		Username:    updUserReq.Username,
		Name:        updUserReq.Name,
		Surname:     updUserReq.Surname,
		PhoneNumber: updUserReq.PhoneNumber,
	}
}
