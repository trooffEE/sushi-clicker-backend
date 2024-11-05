package model

type User struct {
	Id       int64
	Username string `db:"email"`
	Hash     string `db:"hash"`
}
