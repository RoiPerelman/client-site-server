package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"io/ioutil"
	"net/http"
	"github.com/roiperelman/client-site-server/middlewares"
)

func SignupUser(w http.ResponseWriter, r *http.Request) {
	dbStore, ok := r.Context().Value("DBStore").(models.DatabaseStore)
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
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
	id, err := dbStore.InsertUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Id = id
	user.AddToken(middlewares.Secret)

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
