package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
)

func AddContextItem(w http.ResponseWriter, r *http.Request) {
	context := new(models.ContextItem)

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.AddContextTypeItem(context)

	w.Header().Set("content-type", "application/json")
	enc.Encode(context)
}

func DelContextItem(w http.ResponseWriter, r *http.Request) {
	context := new(models.ContextItem)

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	models.DelContextTypeItem(context)

	w.Header().Set("content-type", "application/json")
	enc.Encode(context)
}
