package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type Reasons struct {
	Email        string
	Username     string
	PasswordHash string
}

type UserErrors struct {
	Email            string `json:"email"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Server           string `json:"server"`
	MultipleSections string `json:"multipleSections"`
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
	Errors          UserErrors         `json:"errors"`
	Sections        map[string]Section `json:"sections"`
	IsMulti         bool               `json:"isMulti"`
}

func GetUserById(id int) *User {
	userResults, err := db.Query("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection FROM users WHERE id=?", id)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti)
		if err != nil {
			log.Panic(err)
		}
		user.Sections = GetAllUserIdSections(user.Id)
		return user
	}

	return nil
}

func GetUserByEmail(email string) *User {
	userResults, err := db.Query("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection FROM users WHERE email=?", email)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti)
		if err != nil {
			log.Panic(err)
		}
		user.Sections = GetAllUserIdSections(user.Id)
		return user
	}

	return nil
}

func GetUserByUsername(username string) *User {
	userResults, err := db.Query("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection FROM users WHERE username=?", username)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti)
		if err != nil {
			log.Panic(err)
		}
		user.Sections = GetAllUserIdSections(user.Id)
		return user
	}

	return nil
}

func (user *User) Insert() (int, bool) {
	emailUser := GetUserByEmail(user.Email)
	nameUser := GetUserByUsername(user.Username)

	// if user already exists - add errors indication for client
	if emailUser != nil || nameUser != nil {
		if nameUser != nil {
			user.Errors.Username = "A User with this username exists"
		}
		if emailUser != nil {
			user.Errors.Email = "A User with this email exists"
		}
		return 0, false
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
	//defer insert.Close()
	return int(id), true
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
