package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

const secret = "secret string"

func main() {
	models.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/user/create", createUser).Methods("POST")
	r.HandleFunc("/api/user/authorize", authorizeUser).Methods("GET")
	http.ListenAndServe(":1111", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// create a struct to hold data
	var user models.User
	// Read body to []byte
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Unmarshal []byte to a struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// all methods implicitly add errors to user struct
	if success := user.Insert(); success == true {
		user.AddToken(secret)
		user.SwitchPasswordToPasswordHash()
	} else {
		w.WriteHeader(http.StatusConflict)
	}

	// Marshal the struct to []byte format
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Write the output
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func authorizeUser(	w http.ResponseWriter, r *http.Request) {
	var user models.User

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		user.Errors.Server = err.Error()
		user.IsAuthenticated = false
		w.WriteHeader(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Username = claims["username"].(string)
		user.Email = claims["email"].(string)
		if exists, _ := user.Exists(); exists == true {
			user.IsAuthenticated = true
		} else {
			w.WriteHeader(http.StatusConflict)
			user.IsAuthenticated = false
		}
	} else {
		user.IsAuthenticated = false
		w.WriteHeader(http.StatusUnauthorized)
	}

	// Marshal the struct to []byte format
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Write the output
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
//func createUser2 (w http.ResponseWriter, r *http.Request) {
//	fmt.Println(formatRequest(r))
//
//	var user User // create a struct to hold data
//
//	// create a request.body decoder
//	// which has a method Decode that gets a struct to hold the data
//	dec := json.NewDecoder(r.Body)
//	// create writer encoder
//	// which has a method Encode that gets a struct and writes json response
//	enc := json.NewEncoder(w)
//	err := dec.Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//	w.Header().Set("content-type", "application/json")
//	enc.Encode(user)
//}

//func JsonResponse(response interface{}, w http.ResponseWriter) {
//
//	json, err := json.Marshal(response)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(json)
//}

//err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(user.Password))
//if err != nil {
//	http.Error(w, err.Error(), 500)
//}

//timestamp := claims["timestamp"]
//if t, ok := timestamp.(float64); ok {
//	duration := int64(time.Now().Sub(time.Unix(int64(t), 0)))
//	maximumDuration := int64(3600 * time.Second)
//	fmt.Println(maximumDuration - int64(duration))
//} else {
//	fmt.Println("err")
//	fmt.Println(time.Now().Unix())
//}