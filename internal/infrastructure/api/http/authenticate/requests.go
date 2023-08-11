package authenticate

type SignUpRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleCode    string `json:"role"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
