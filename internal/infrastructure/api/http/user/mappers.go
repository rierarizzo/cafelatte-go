package user

import (
	"github.com/rierarizzo/cafelatte/internal/domain/user"
	"strings"
)

func FromSignUpReqToUser(req SignUpRequest) user.User {
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

func FromUserToUserRes(user user.User) Response {
	return Response{
		ID:           user.ID,
		CompleteName: strings.Join([]string{user.Name, user.Surname}, " "),
		PhoneNumber:  user.PhoneNumber,
		Username:     user.Username,
		Email:        user.Email,
		Role:         user.RoleCode,
	}
}

func FromUserSliceToUserResSlice(users []user.User) []Response {
	var res = make([]Response, 0)
	for _, v := range users {
		res = append(res, FromUserToUserRes(v))
	}

	return res
}

func FromUpdateUserReqToUser(updUserReq UpdateUserRequest) user.User {
	return user.User{
		Username:    updUserReq.Username,
		Name:        updUserReq.Name,
		Surname:     updUserReq.Surname,
		PhoneNumber: updUserReq.PhoneNumber,
	}
}
