package domain

type AuthorizedUser struct {
	User        User
	AccessToken string
}
