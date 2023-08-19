package domain

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

func (u *User) SetPassword(password string) {
	u.Password = password
}

// AuthorizedUser represents an user with authorization tokens
type AuthorizedUser struct {
	User        User
	AccessToken string
}
