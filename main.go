package main

import (
	"con-system/dao/mysql"
	"con-system/pkg/sonyflake"
	"con-system/router"
	"con-system/settings"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	id   int
	name string
	age  int
}

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 连接数据库
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	fmt.Println("数据库连接成功！")
	defer mysql.Close() // 程序退出关闭数据库连接

	if err := sonyflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	id, err := sonyflake.GetId()
	if err != nil {
		fmt.Printf("GetId snowflake failed, err:%v\n", err)
		return
	}
	fmt.Println("id:", id)

	// 注册路由
	r := router.SetupRouter()
	err = r.Run()
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
