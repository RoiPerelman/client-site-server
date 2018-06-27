package handlers

import (
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"io/ioutil"
	"encoding/json"
)

const secret = "secret string"

func SignupUser(w http.ResponseWriter, r *http.Request) {
	// create a struct to hold data
	var user models.User
	// Read body to []byte
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Unmarshal []byte to a struct
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.SwitchPasswordToPasswordHash()
	// insert method implicitly adds errors to user struct
	if id, success := user.Insert(); success == true {
		user.Id = id
		user.AddToken(secret)
	} else {
		w.WriteHeader(http.StatusConflict)
	}

	// Marshal the struct to []byte format
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the output
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}