package authenticator

type AuthenticatedResponse struct {
	User struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	} `json:"user"`
	AccessToken string `json:"accessToken"`
}
