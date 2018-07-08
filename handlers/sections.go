package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
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
	var section models.Section

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user := r.Context().Value("User"); user != nil {
		sectionsId := models.AddSection(user.(models.User).Id, section)
		section = models.GetUserSectionBySectionsId(sectionsId)
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}

func DelSectionUser(w http.ResponseWriter, r *http.Request) {
	var section models.Section

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user := r.Context().Value("User"); user != nil {
		if section.SectionId != user.(models.User).DefaultSection {
			models.DelSection(user.(models.User).Id, section)
		} else {
			http.Error(w, "Trying to remove default section!", http.StatusForbidden)
			return
		}
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}
