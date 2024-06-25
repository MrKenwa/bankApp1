package utils

import "golang.org/x/crypto/bcrypt"

func IsPasswordCorrect(passwd, hashedPasswd []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPasswd, passwd); err != nil {
		return false
	}
	return true
}
