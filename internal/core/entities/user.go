package entities

import "strings"

type User struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	PhoneNumber string
	Email       string
	Password    string
	/* A: Administrador, E: Empleado, C: Cliente */
	RoleCode     string
	Addresses    []Address
	PaymentCards []PaymentCard
}

func (u *User) IsValidUser() bool {
	return u.isValidRole() && u.isValidPhoneNumber() && u.isValidEmail()
}

func (u *User) isValidRole() bool {
	return u.RoleCode == "A" || u.RoleCode == "E" || u.RoleCode == "C"
}

func (u *User) isValidPhoneNumber() bool {
	return len(u.PhoneNumber) == 10
}

func (u *User) isValidEmail() bool {
	return strings.Contains(u.Email, "@")
}

func (u *User) SetPassword(password string) {
	u.Password = password
}
