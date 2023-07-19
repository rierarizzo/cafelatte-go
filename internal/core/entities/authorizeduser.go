package entities

type AuthorizedUser struct {
	User        User
	AccessToken string
}
