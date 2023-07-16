package entities

type AuthorizedUser struct {
	UserInfo    User
	AccessToken string
}
