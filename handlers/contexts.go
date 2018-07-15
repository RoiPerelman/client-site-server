package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
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

	section = models.GetUserSectionBySectionsId(context.SectionsId)

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

	section = models.GetUserSectionBySectionsId(context.SectionsId)

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}
