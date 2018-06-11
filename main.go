package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"github.com/roiperelman/client-site-server/models"
)

type UserErrors struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

type User struct {
	Email        string     `json:"email"`
	Username     string     `json:"username"`
	Password     string     `json:"password"`
	PasswordHash string     `json:"passwordHash"`
	Token        string     `json:"token"`
	Errors       UserErrors `json:"errors"`
}

func main() {
	models.InitDB()
	defer models.Close()

	r := mux.NewRouter()

	r.HandleFunc("/api/user/create", createUser).Methods("POST")
	http.ListenAndServe(":1111", r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatRequest(r))
	// create a struct to hold data
	var user User
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.Username,
		"email":     user.Email,
		"timestamp": time.Now().Unix(),
	})

	secret := []byte("password")
	tokenString, err := token.SignedString(secret)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	user.Token = tokenString

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.PasswordHash = string(passwordHash)

	//err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(user.Password))
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//}
	user.Password = ""

	insert, err := db.Query(`
		INSERT INTO users (email,username,passwordHash)
		VALUES ( 'r@g.com', 'roi', 'hashedpassword' )
	`)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	//token2, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	// Don't forget to validate the alg is what you expect:
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	//	}
	//
	//	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	//	return secret, nil
	//})
	//
	//if claims, ok := token2.Claims.(jwt.MapClaims); ok && token2.Valid {
	//	fmt.Println(claims["username"], claims["password"], claims["timestamp"])
	//} else {
	//	fmt.Println(err)
	//}

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

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	//if r.Method == "POST" {
	//	var buf bytes.Buffer
	//	io.TeeReader(r.Body, &buf)
	//	body, _ := ioutil.ReadAll(&buf)
	//	if len(body) > 0 {
	//		request = append(request, fmt.Sprintf("body: %v", string(body)))
	//	}
	//} else if r.Method == "GET" {
	//	queryParams := r.URL.Query()
	//	request = append(request, "query:")
	//	for key, values := range queryParams {
	//		key = strings.ToLower(key)
	//		for _, value := range values {
	//			request = append(request, fmt.Sprintf("%v: %v", key, value))
	//		}
	//	}
	//}
	// Return the request as a string
	return strings.Join(request, "\n")
}

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
