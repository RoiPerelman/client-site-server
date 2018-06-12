package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() () {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/client_site_server_development")
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}

//func Close() {
//	db.Close()
//}