package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
	"fmt"
)

type UserErrors struct {
	Email    string `json:"email"`
	Username     string `json:"name"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

type User struct {
	Id           int        `json:"id"`
	Email        string     `json:"email"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	PasswordHash string     `json:"passwordHash"`
	Token        string     `json:"token"`
	Errors       UserErrors `json:"errors"`
}

func UserExists(user *User) bool {
	results, err := db.Query("SELECT email, username FROM users WHERE email='roiperelman@gmail.com'")
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
			user.Errors.Email = "u user with this email exists"
			response = true
		}
		if userName == user.Username {
			user.Errors.Username = "a user with this username exists"
			response = true
		}
		return response
	}
	return false
}

func InsertUser(user *User) {

	if UserExists(user) {
		return
	} else {
		str := fmt.Sprintf(
			`INSERT INTO users (email,username,passwordHash)
		VALUES ( '%v', '%v', '%v' )`, user.Email, user.Username, user.PasswordHash)
		insert, err := db.Query(str)

		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()
		return
	}
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
	return err
}

func (user *User) SwitchPasswordToPasswordHash() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.PasswordHash = string(passwordHash)
	user.Password = ""
	return err
}
