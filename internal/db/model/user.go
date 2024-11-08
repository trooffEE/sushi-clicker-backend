package model

type User struct {
	Id    int64
	Email string `db:"email"`
	Hash  string `db:"hash"`
}
