package authenticate

type AuthorizedUserResponse struct {
	User        LoggedUserResponse `json:"user"`
	AccessToken string             `json:"accessToken"`
}

type LoggedUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
