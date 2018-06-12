package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/roiperelman/client-site-server/models"
)

//type UserErrors struct {
//	Email    string `json:"email"`
//	Name     string `json:"name"`
//	Password string `json:"password"`
//	Server   string `json:"server"`
//}
//
//type User struct {
//	Email        string     `json:"email"`
//	Username     string     `json:"username"`
//	Password     string     `json:"password"`
//	PasswordHash string     `json:"passwordHash"`
//	Token        string     `json:"token"`
//	Errors       UserErrors `json:"errors"`
//}

func main() {
	models.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/api/user/create", createUser).Methods("POST")
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

	models.InsertUser(&user)

	user.AddToken("secret string")
	user.SwitchPasswordToPasswordHash()

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