package mysql

import (
	"con-system/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "Kphilleani"

// CheckUserExist 检查用户是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) error {
	// 执行SQl语句
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	password := encryptPassword(user.Password)
	_, err := db.Exec(sqlStr, user.UserId, user.UserName, password)
	if err != nil {
		return err
	}
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func Login(user models.User) (err error) {
	var u models.User
	sqlStr := `select username, password from user where username = ?`
	err = db.Get(&u, sqlStr, user.UserName)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return err
	}
	if encryptPassword(user.Password) != u.Password {
		return errors.New("密码错误")
	}
	return
}
