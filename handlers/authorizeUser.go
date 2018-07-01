package handlers

import (
	"net/http"
	//"github.com/roiperelman/client-site-server/models"
	"encoding/json"
)

func AuthorizeUser(	w http.ResponseWriter, r *http.Request) {
	//Get data from context
	if user := r.Context().Value("User"); user != nil {
		w.WriteHeader(http.StatusOK)
		output, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Write the output
		w.Header().Set("content-type", "application/json")
		w.Write(output)
	} else {
		http.Error(w, "authrize user failed", http.StatusInternalServerError)
		return
	}

}