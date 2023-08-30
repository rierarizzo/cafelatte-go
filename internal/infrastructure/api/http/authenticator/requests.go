package authenticator

type UserSignup struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleCode    string `json:"role"`
}

type UserSignin struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
