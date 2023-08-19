package authenticator

type AuthorizedUserResponse struct {
	User        LoggedUserResponse `json:"usermanager"`
	AccessToken string             `json:"accessToken"`
}

type LoggedUserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
