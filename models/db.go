package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/roiperelman/client-site-server/utils"
)

type DBUserJSStore interface {
	UpdateJSCode(id int, jsCode string) error
}
type DBUserStore interface {
	InsertUser(user *User) (int, error)
	GetUserById(id int) *User
	GetUserByEmail(email string) *User
	GetUserByUsername(username string) *User
	DBUserJSStore
}

type DBSectionStore interface {
	GetAllUserIdSections(userId int) map[string]Section
	GetUserSectionBySectionsId(sectionsId int) Section
	AddSection(userId int, section Section) int
	DelSection(id int, section Section)
	UpdateIsMultipleSectionFeature(id int, isMulti bool)
}

type DBContextStore interface {
	GetContextsBySectionsId(sectionsIdentifier int) Contexts
	AddContextTypeItem(contextItem *ContextItem)
	DelContextTypeItem(contextItem *ContextItem)
}
// DatabaseStore is a DB abstraction that hold all db methods
type DatabaseStore interface {
	DBUserStore
	DBSectionStore
	DBContextStore
}

type DB struct {
	*sql.DB
}

func InitDB() (*DB, error){
	var err error

	address := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		utils.GetEnv("MYSQL_USERNAME", "root"),
		utils.GetEnv("MYSQL_PASSWORD", "root"),
		utils.GetEnv("MYSQL_HOST", "localhost"),
		utils.GetEnv("MYSQL_PORT", "3306"),
		utils.GetEnv("MYSQL_DATABASE", "dyrp_dev"))

	db, err := sql.Open("mysql", address)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}