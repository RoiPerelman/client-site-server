package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(models.DBStore); ok {
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

		dbUser := dbStore.GetUserByEmail(user.Email)
		if dbUser != nil {
			err = bcrypt.CompareHashAndPassword([]byte(dbUser.PasswordHash), []byte(user.Password))
			if err != nil {
				http.Error(w, "Error: Invalid Email/password combination!", http.StatusUnauthorized)
				return
			} else {
				dbUser.AddToken(secret)
				dbUser.IsAuthenticated = true
			}
		} else {
			http.Error(w, "Error: Invalid Email/password combination!", http.StatusUnauthorized)
			return
		}

		//user.Password = ""
		w.Header().Set("content-type", "application/json")
		enc.Encode(dbUser)
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}

}
