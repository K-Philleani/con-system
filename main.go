package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	id int
	name string
	age int
}

func initMysql() (err error){
	dsn := "root:123456@tcp(124.70.71.78:3306)/cons"
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

func queryRowDemo() {
	sqlStr := "select id, name, age from user where id = ?"
	var u User
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d name: %s age:%d\n", u.id, u.name, u.age)
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()
	fmt.Println("数据库连接成功")
	queryRowDemo()
}
