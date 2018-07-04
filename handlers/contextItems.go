package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
)

func AddContextItem(w http.ResponseWriter, r *http.Request) {
	context := new(models.ContextItem)
	var section models.Section

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.AddContextTypeItem(context)

	if user := r.Context().Value("User"); user != nil {
		section = models.GetUserIdSectionBySectionId(user.(models.User).Id, context.SectionId)
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}

func DelContextItem(w http.ResponseWriter, r *http.Request) {
	context := new(models.ContextItem)
	var section models.Section

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.DelContextTypeItem(context)

	if user := r.Context().Value("User"); user != nil {
		section = models.GetUserIdSectionBySectionId(user.(models.User).Id, context.SectionId)
	} else {
		http.Error(w, "authorize user failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}
