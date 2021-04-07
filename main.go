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

func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id: %d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func insertRowDemo() {
	sqlStr := "insert into user (name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "宋元策", 25)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsertid failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is: %d\n", theId)
}

func updateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 30, 4)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 5)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsaffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

func main() {
	if err := initMysql(); err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()
	fmt.Println("数据库连接成功")
	queryRowDemo()
	fmt.Println("==========================")
	queryMultiRowDemo()
	fmt.Println("==========================")
	//insertRowDemo()
	fmt.Println("==========================")
	updateRowDemo()
	fmt.Println("==========================")
	deleteRowDemo()
}
