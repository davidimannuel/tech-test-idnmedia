package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryptEncrypt(secret string) string {
	res, _ := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	return string(res)
}

func BcryptCompare(hash, secret string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret)) == nil
}
