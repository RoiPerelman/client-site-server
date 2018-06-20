package models

import (
	//_ "github.com/go-sql-driver/mysql"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/roiperelman/client-site-server/utils"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", utils.GetEnv("DATABASE_URL", "dbname=mylocaldb sslmode=disable"))

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}
