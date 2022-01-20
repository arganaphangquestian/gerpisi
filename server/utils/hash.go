package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHash(plainText string) (*string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	if err != nil {
		return nil, err
	}
	res := string(bytes)
	return &res, err
}

func CompareHash(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
