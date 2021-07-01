package DBOP

import "database/sql"

func InitDB() (*sql.DB, error) {
	dsn := "root:123456@tcp(124.70.71.78:3306)/cons"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
