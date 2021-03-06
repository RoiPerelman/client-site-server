package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"encoding/json"
	"database/sql"
)

type UserErrors struct {
	Email            string `json:"email"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Login            string `json:"login"`
	MultipleSections string `json:"multipleSections"`
}

func (e UserErrors) Error() string {
	output, err := json.Marshal(e)
	if err != nil {
		return e.Error()
	}
	return string(output)
}

type User struct {
	Id              int                `json:"id"`
	Email           string             `json:"email"`
	Username        string             `json:"username"`
	Password        string             `json:"password"`
	DefaultSection  string             `json:"defaultSection"`
	PasswordHash    string             `json:"-"`
	Token           string             `json:"token"`
	IsAuthenticated bool               `json:"isAuthenticated"`
	Sections        map[string]Section `json:"sections"`
	IsMulti         bool               `json:"isMulti"`
	JSCode          string             `json:"jsCode"`
}

func (db *DB) GetUserById(id int) *User {
	user := new(User)
	err := db.QueryRow("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection, JSCode FROM users WHERE id=?", id).
		Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti, &user.JSCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Panic(err)
	}
	user.Sections = db.GetAllUserIdSections(user.Id)
	return user
}

func (db *DB) GetUserByEmail(email string) *User {
	userResults, err := db.Query("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection, JSCode FROM users WHERE email=?", email)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti, &user.JSCode)
		if err != nil {
			log.Panic(err)
		}
		user.Sections = db.GetAllUserIdSections(user.Id)
		return user
	}

	return nil
}

func (db *DB) GetUserByUsername(username string) *User {
	userResults, err := db.Query("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection, JSCode FROM users WHERE username=?", username)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti, &user.JSCode)
		if err != nil {
			log.Panic(err)
		}
		user.Sections = db.GetAllUserIdSections(user.Id)
		return user
	}

	return nil
}

func (db *DB) UpdateJSCode(id int, jsCode string) error {
	_, err := db.Exec(`UPDATE users SET jsCode=? WHERE id=?`, jsCode, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertUser(user *User) (int, error) {
	userErrors := new(UserErrors)
	emailUser := db.GetUserByEmail(user.Email)
	nameUser := db.GetUserByUsername(user.Username)

	// if user already exists - add errors indication for client
	if emailUser != nil || nameUser != nil {
		if nameUser != nil {
			userErrors.Username = "A User with this username exists"
		}
		if emailUser != nil {
			userErrors.Email = "A User with this email exists"
		}
		return 0, userErrors
	}

	// create user in database
	insert, err := db.Exec(
		`INSERT INTO users (email, username, passwordHash, DefaultSection)
			VALUES (?, ?, ?, ?)`, user.Email, user.Username, user.PasswordHash, user.DefaultSection)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	id, err := insert.LastInsertId()
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	//create section in database
	insert, err = db.Exec(
		`INSERT INTO sections (userId, sectionId)
			VALUES (?, ?)`, id, user.DefaultSection)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	return int(id), nil
}

func (user *User) AddToken(secret string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.Id,
		"username":  user.Username,
		"email":     user.Email,
		"timestamp": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	user.Token = tokenString
	user.IsAuthenticated = true
	return err
}

func (user *User) SwitchPasswordToPasswordHash() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.PasswordHash = string(passwordHash)
	user.Password = ""
	return err
}
