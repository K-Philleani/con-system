package logic

import (
	"con-system/dao/mysql"
	"con-system/pkg/sonyflake"
)

func SignUp() {
	// 1.判断用户是否存在
	mysql.QueryUserById()
	// 2.生成uid
	sonyflake.GetId()
	// 3.保存进数据库
	mysql.InsertUser()
}
