package user

import (
	domain "github.com/rierarizzo/cafelatte/internal/domain/errors"
	"github.com/rierarizzo/cafelatte/pkg/utils/crypt"
)

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	PhoneNumber string
	Email       string
	Password    string
	/* A: Admin, E: Employee, C: Client */
	RoleCode string
}

func (u *User) HashPassword() *domain.AppError {
	hashed, appErr := crypt.HashText(u.Password)
	if appErr != nil {
		return appErr
	}
	u.Password = hashed

	return nil
}