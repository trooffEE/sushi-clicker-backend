package user

import "errors"

var (
	IsAlreadyRegistered  = errors.New("User is already registered")
	IncorrectCredentials = errors.New("Incorrect credentials")
)
