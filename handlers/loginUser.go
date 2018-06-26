package handlers

import (
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User // create a struct to hold data

	// create a request.body decoder
	// which has a method Decode that gets a struct to hold the data
	dec := json.NewDecoder(r.Body)
	// create writer encoder
	// which has a method Encode that gets a struct and writes json response
	enc := json.NewEncoder(w)
	err := dec.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbUser := models.GetUserByEmail(user.Email)
	if dbUser != nil {
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
		if err != nil {
			user.IsAuthenticated = false
			user.Errors.Server = "Error: Invalid Email/password combination!"
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			user.Username = dbUser.Username
			user.Id = dbUser.Id
			user.DefaultSection = dbUser.DefaultSection
			user.AddToken(secret)
			user.IsAuthenticated = true
		}
	} else {
		user.IsAuthenticated = false
		user.Errors.Server = "Error: Invalid Email/password combination!"
		w.WriteHeader(http.StatusUnauthorized)
	}

	//user.Password = ""
	w.Header().Set("content-type", "application/json")
	enc.Encode(user)
}
