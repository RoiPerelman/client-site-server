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
	//db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/client_site_server_development")
	fmt.Printf("DATABASE_URL CONNECTION TO " + os.Getenv("DATABASE_URL"))
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
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