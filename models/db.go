package models

import (
	"database/sql"
	"fmt"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/roiperelman/client-site-server/utils"
	"net/http"
)

// DatabaseStore is a DB abstraction that hold all db methods
type DatabaseStore interface {
	// users
	InsertUser(user *User) (int, error)
	GetUserById(id int) *User
	GetUserByEmail(email string) *User
	GetUserByUsername(username string) *User
	UpdateJSCode(id int, jsCode string) error
	// sections
	GetAllUserIdSections(userId int) map[string]Section
	GetUserSectionBySectionsId(sectionsId int) Section
	AddSection(userId int, section Section) int
	DelSection(id int, section Section)
	UpdateIsMultipleSectionFeature(id int, isMulti bool)
	// contexts
	GetContextsBySectionsId(sectionsIdentifier int) Contexts
	AddContextTypeItem(contextItem *ContextItem)
	DelContextTypeItem(contextItem *ContextItem)
}

type DBStore struct {
	DatabaseStore
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

func (dbStore *DBStore) DBStoreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "DBStore", dbStore)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}