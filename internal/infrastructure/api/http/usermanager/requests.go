package usermanager

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phone"`
}
