package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
	"log"
)

type Reasons struct {
	Email        string
	Username     string
	PasswordHash string
}

type UserErrors struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

type User struct {
	Id              int        `json:"-"`
	Email           string     `json:"email"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	SectionId       string     `json:"sectionId"`
	PasswordHash    string     `json:"-"`
	Token           string     `json:"token"`
	IsAuthenticated bool       `json:"isAuthenticated"`
	Errors          UserErrors `json:"errors"`
}

func GetUserByEmail(email string) *User {
	results, err := db.Query("SELECT email, username, passwordHash, sectionId FROM users WHERE email=$1", email)
	if err != nil {
		log.Panic(err)
	}
	defer results.Close()
	found := results.Next()
	if found {
		user := new(User)
		err = results.Scan(&user.Email, &user.Username, &user.PasswordHash, &user.SectionId)
		return user
	}
	return nil
}

func GetUserByUsername(email string) *User {
	results, err := db.Query("SELECT email, username, passwordHash, sectionId FROM users WHERE username=$1", email)
	if err != nil {
		log.Panic(err)
	}
	defer results.Close()
	found := results.Next()
	if found {
		user := new(User)
		err = results.Scan(&user.Email, &user.Username, &user.PasswordHash, &user.SectionId)
		return user
	}
	return nil
}

func (user *User) Insert() bool {
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
		return false
	}

	// create user in database
	str := fmt.Sprintf(
		`INSERT INTO users (email,username,passwordHash,SectionId)
			VALUES ( '%v', '%v', '%v', '%v' )`, user.Email, user.Username, user.PasswordHash, user.SectionId)
	insert, err := db.Query(str)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	defer insert.Close()
	return true
}

func (user *User) AddToken(secret string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
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
