package domain

type User struct {
	Id          int
	Username    string
	Name        string
	Surname     string
	PhoneNumber string
	Email       string
	Password    string
	/* A: Admin, E: Employee, C: Client */
	RoleCode string
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

type AuthenticatedUser struct {
	User        User
	AccessToken string
}
