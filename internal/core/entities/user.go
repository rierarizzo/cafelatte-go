package entities

import "strings"

type User struct {
	ID          int
	Name        string
	Surname     string
	PhoneNumber string
	Email       string
	Password    string
	/* A: Administrador, E: Empleado, C: Cliente */
	Role      string
	Addresses []Address
	/* E: Efectivo, T: Tarjeta */
	PaymentMethod string
	Card          PaymentCard
}

func (u *User) IsValidUser() bool {
	return u.isValidRole() && u.isValidPhoneNumber() && u.isValidEmail() && u.isValidPaymentMethod()
}

func (u *User) isValidRole() bool {
	return u.Role == "A" || u.Role == "E" || u.Role == "C"
}

func (u *User) isValidPhoneNumber() bool {
	return len(u.PhoneNumber) == 10
}

func (u *User) isValidEmail() bool {
	return strings.Contains(u.Email, "@")
}

func (u *User) isValidPaymentMethod() bool {
	return u.PaymentMethod == "E" || u.PaymentMethod == "T"
}
