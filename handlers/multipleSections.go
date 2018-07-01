package handlers

import (
"net/http"
"encoding/json"
	"github.com/roiperelman/client-site-server/models"
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

	if user := r.Context().Value("User"); user != nil {
		models.UpdateIsMultipleSectionFeature(user.(models.User).Id, isMulti)
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(isMulti)
}

func AddSectionUser(w http.ResponseWriter, r *http.Request) {
	var section float64 // create a struct to hold data

	// create a request.body decoder
	// which has a method Decode that gets a struct to hold the data
	dec := json.NewDecoder(r.Body)
	// create writer encoder
	// which has a method Encode that gets a struct and writes json response
	enc := json.NewEncoder(w)
	err := dec.Decode(&section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user := r.Context().Value("User"); user != nil {
		models.AddSection(user.(models.User).Id, int(section))
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}

func DelSectionUser(w http.ResponseWriter, r *http.Request) {
	var section float64 // create a struct to hold data

	// create a request.body decoder
	// which has a method Decode that gets a struct to hold the data
	dec := json.NewDecoder(r.Body)
	// create writer encoder
	// which has a method Encode that gets a struct and writes json response
	enc := json.NewEncoder(w)
	err := dec.Decode(&section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user := r.Context().Value("User"); user != nil {
		models.DelSection(user.(models.User).Id, int(section))
	} else {
		http.Error(w, "authrize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}