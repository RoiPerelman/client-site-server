package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ExistencReasons struct {
	Email string
	Username string
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
	Password        string     `json:"-"`
	PasswordHash    string     `json:"-"`
	Token           string     `json:"token"`
	IsAuthenticated bool       `json:"isAuthenticated"`
	Errors          UserErrors `json:"errors"`
}

func (user *User) Exists() (bool, ExistencReasons) {
	var existenceReasons ExistencReasons
	results, err := db.Query(
		"SELECT email, username FROM users WHERE email=? AND username=?",
		user.Email, user.Username)
	if err != nil {
		panic(err.Error())
	}
	defer results.Close()

	found := results.Next()

	if found {
		var response = false
		var email, userName string
		err = results.Scan(&email, &userName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		if email == user.Email {
			existenceReasons.Email = "A User with this email exists"
			response = true
		}
		if userName == user.Username {
			existenceReasons.Username = "A User with this username exists"
			response = true
		}
		return response, existenceReasons
	}
	return false, existenceReasons
}

func (user *User) Insert() bool {
	if exists, existenceReasons := user.Exists(); exists == true {
		user.Errors.Email = existenceReasons.Email
		user.Errors.Username = existenceReasons.Username
		return false
	}
	str := fmt.Sprintf(
		`INSERT INTO users (email,username,passwordHash)
			VALUES ( '%v', '%v', '%v' )`, user.Email, user.Username, user.PasswordHash)
	insert, err := db.Query(str)
	if err != nil {
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
