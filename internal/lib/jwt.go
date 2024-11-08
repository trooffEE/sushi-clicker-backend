package lib

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWTToken(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	fmt.Println(string(hash))
	return string(hash), nil
}

//func getHashPassword(password string) (string, error) {
//	bytePassword := []byte(password)
//	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
//	if err != nil {
//		return "", err
//	}
//	return string(hash), nil
//}
