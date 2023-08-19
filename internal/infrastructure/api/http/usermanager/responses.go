package usermanager

type Response struct {
	ID           int    `json:"id"`
	CompleteName string `json:"completeName"`
	Username     string `json:"username"`
	PhoneNumber  string `json:"phoneNumber"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}
