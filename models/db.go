package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var db *sql.DB
var err error

func InitDB() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/client_site_server_development")

	if err != nil {
		log.Panic(err)
		panic(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func Close() {
	db.Close()
}