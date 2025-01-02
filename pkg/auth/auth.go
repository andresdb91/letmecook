package auth

import auth "github.com/andresdb91/letmecook/internal/business/auth"

func CheckValidUser(token string) (bool, error) {
	return auth.CheckValidUser(token)
}
