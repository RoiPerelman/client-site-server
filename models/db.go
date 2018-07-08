package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/roiperelman/client-site-server/utils"
	"log"
)

var db *sql.DB

func InitDB() {
	var err error

	address := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		utils.GetEnv("MYSQL_USERNAME", "root"),
		utils.GetEnv("MYSQL_PASSWORD", "root"),
		utils.GetEnv("MYSQL_HOST", "localhost"),
		utils.GetEnv("MYSQL_PORT", "3306"),
		utils.GetEnv("MYSQL_DATABASE", "dyrp_dev"))

	db, err = sql.Open("mysql", address)

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}
}
