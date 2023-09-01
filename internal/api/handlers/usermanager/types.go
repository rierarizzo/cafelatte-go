package usermanager

type UserUpdate struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
}

type UserResponse struct {
	Id           int    `json:"id"`
	CompleteName string `json:"completeName"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}
