package handlers

import (
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"encoding/json"
)

func UpdateJSCode(w http.ResponseWriter, r *http.Request) {
	var payload struct{
		JsCode string `json:"jsCode"`
	}

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	err := dec.Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if id, ok := r.Context().Value("UserId").(int); ok {
		err = models.UpdateJSCode(id, payload.JsCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "update JS code failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(payload)
}