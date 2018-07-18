package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/roiperelman/client-site-server/middlewares"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	dbStore, ok := r.Context().Value("DBStore").(models.DatabaseStore)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// create a struct to hold data
	var user models.User
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

	dbUser := dbStore.GetUserByEmail(user.Email)
	if dbUser == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	dbUser.AddToken(middlewares.Secret)
	dbUser.IsAuthenticated = true

	w.Header().Set("content-type", "application/json")
	enc.Encode(dbUser)
}
