package models

import (
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"os"
	"fmt"
)

var db *sql.DB

func InitDB() () {
	var err error
	fmt.Printf("DATABASE_URL CONNECTION TO " + os.Getenv("DATABASE_URL"))
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/client_site_server_development")
	//db, err = sql.Open("postgres", "postgres://ttfhkthanlqreb:632e0cd045e0640c42395f3359b998ff8f0105615ca9bbdfbbf49cc8f36fb6cf@ec2-107-21-255-2.compute-1.amazonaws.com:5432/dbpda6g26nheoc")
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