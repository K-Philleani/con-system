package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initMysql() (err error){
	dsn := "root:123456@tcp(124.70.71.78:3306)/db1"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()
	fmt.Println("数据库连接成功")
}
