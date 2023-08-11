package authenticate

import (
	"github.com/rierarizzo/cafelatte/internal/domain/user"
)

type AuthorizedUser struct {
	User        user.User
	AccessToken string
}
