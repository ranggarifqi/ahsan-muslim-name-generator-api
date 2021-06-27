package auth

import "github.com/ranggarifqi/ahsan-muslim-name-generator-api/src/user"

type SignInResult struct {
	user.UserWithoutPassword
	Token string `json:"token"`
}
