package logic

import (
	"con-system/dao/mysql"
	"con-system/models"
	"con-system/pkg/sonyflake"
)

func SignUp(p *models.ParamSignup) (err error) {
	// 1.判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	// 2.生成uid
	userId, err := sonyflake.GetId()
	// 3.构造一个User实例
	u := models.User{
		UserId:   userId,
		UserName: p.Username,
		Password: p.Password,
	}
	// 4.保存进数据库
	if err = mysql.InsertUser(&u); err != nil {
		return
	}
	return
}

func Login(p models.ParamLogin) (err error) {
	user := models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}
