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
	Id              int        `json:"id"`
	Email           string     `json:"email"`
	Username        string     `json:"username"`
	Password        string     `json:"password"`
	DefaultSection  string     `json:"defaultSection"`
	PasswordHash    string     `json:"-"`
	Token           string     `json:"token"`
	IsAuthenticated bool       `json:"isAuthenticated"`
	Errors          UserErrors `json:"errors"`
	Sections        []Section   `json:"sections"`
	IsMulti         bool       `json:"isMulti"`
}

func GetUserByEmail(email string) *User {
	// get User info
	query := fmt.Sprintf("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection FROM users WHERE email='%v'", email)
	userResults, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	defer userResults.Close()
	found := userResults.Next()
	if found {
		user := new(User)
		user.Sections = make([]Section, 0)
		err = userResults.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti)
		if err != nil {
			log.Panic(err)
		}

		// get Sections info
		query := fmt.Sprintf("Select id, sectionId FROM sections WHERE sections.userId=%v", user.Id)
		sectionResults, err := db.Query(query)
		if err != nil {
			log.Panic(err)
		}
		defer sectionResults.Close()

		for sectionResults.Next() {
			section := new(Section)
			section.Contexts.ProductContext = make([]string, 0)
			section.Contexts.CartContext = make([]string, 0)
			section.Contexts.CategoryContext = make([]string, 0)
			err := sectionResults.Scan(&section.Id, &section.SectionId)
			if err != nil {
				log.Fatal(err)
			}

			query := fmt.Sprintf("Select type, item FROM contexts WHERE contexts.sectionsId=%v", section.Id)
			contextResults, err := db.Query(query)
			if err != nil {
				log.Panic(err)
			}
			defer contextResults.Close()

			for contextResults.Next() {
				contextItem := new(ContextItem)
				err := contextResults.Scan(&contextItem.ContextType, &contextItem.Item)
				if err != nil {
					log.Fatal(err)
				}

				switch contextType := contextItem.ContextType; contextType {
				case "PRODUCT":
					section.Contexts.ProductContext = append(section.Contexts.ProductContext, contextItem.Item)
				case "CART":
					section.Contexts.CartContext = append(section.Contexts.CartContext, contextItem.Item)
				case "CATEGORY":
					section.Contexts.CategoryContext = append(section.Contexts.CategoryContext, contextItem.Item)
				}
			}

			user.Sections = append(user.Sections, *section)
		}
		return user
	}

	return nil
}

func GetUserByUsername(username string) *User {
	query := fmt.Sprintf("SELECT id, email, username, passwordHash, DefaultSection, isMultipleSection FROM users WHERE username='%v'", username)
	results, err := db.Query(query)
	if err != nil {
		log.Panic(err)
	}
	defer results.Close()
	found := results.Next()
	if found {
		user := new(User)
		err = results.Scan(&user.Id, &user.Email, &user.Username, &user.PasswordHash, &user.DefaultSection, &user.IsMulti)
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
	query := fmt.Sprintf(
		`INSERT INTO users (email, username, passwordHash, DefaultSection)
			VALUES ('%v', '%v', '%v', '%v')`, user.Email, user.Username, user.PasswordHash, user.DefaultSection)
	insert, err := db.Exec(query)
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
	query = fmt.Sprintf(
		`INSERT INTO sections (userId, sectionId)
			VALUES ('%v', '%v')`, id, user.DefaultSection)
	insert, err = db.Exec(query)
	if err != nil {
		fmt.Printf("insert err %v\n", err.Error())
		panic(err.Error())
	}
	//defer insert.Close()
	return int(id), true
}

func (user *User) AddToken(secret string) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
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
