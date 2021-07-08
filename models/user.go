package models

type User struct {
	UserId   uint64 `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
}
