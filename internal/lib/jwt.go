package lib

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var (
	InvalidTokenError = errors.New("invalid token")
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
var secretKeyRefresh = []byte(os.Getenv("JWT_SECRET_KEY_REFRESH"))

func GenerateJwtRefreshToken(email string) (string, time.Time, error) {
	exp := time.Now().Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   exp.Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenSigned, err := token.SignedString(secretKeyRefresh)

	return tokenSigned, exp, err
}

func GenerateJwtAccessToken(email, sugar string) (string, error) {
	exp := time.Now().Add(time.Minute * 1) // TODO replace 1 with 15 or something
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"sugar": sugar,
		"exp":   exp,
		"iat":   time.Now().Unix(),
	})

	tokenSigned, err := token.SignedString(secretKey)

	return tokenSigned, err
}

func ValidateJwtAccessToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, InvalidTokenError
	}

	return token, nil
}

func ValidateJwtRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyRefresh, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, InvalidTokenError
	}

	return token, nil
}
