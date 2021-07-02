package main

import (
	"con-system/pkg/sonyflake"
	"con-system/router"
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
