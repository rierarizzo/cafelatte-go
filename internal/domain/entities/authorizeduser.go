package entities

type AuthorizedUser struct {
	User        User
	AccessToken string
}

func NewAuthorizedUser(user User, accessToken string) *AuthorizedUser {
	return &AuthorizedUser{
		User:        user,
		AccessToken: accessToken,
	}
}
