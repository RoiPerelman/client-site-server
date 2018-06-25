package handlers

import (
"net/http"
"encoding/json"
)

func MultipleSectionsUser(w http.ResponseWriter, r *http.Request) {
	var isMulti bool // create a struct to hold data

	// create a request.body decoder
	// which has a method Decode that gets a struct to hold the data
	dec := json.NewDecoder(r.Body)
	// create writer encoder
	// which has a method Encode that gets a struct and writes json response
	enc := json.NewEncoder(w)
	err := dec.Decode(&isMulti)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//user.Password = ""
	w.Header().Set("content-type", "application/json")
	enc.Encode([]byte("aaaaaa"))
}
