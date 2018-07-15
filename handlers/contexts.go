package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
)

func AddContextItem(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(models.DatabaseStore); ok {
		context := new(models.ContextItem)
		var section models.Section

		dec := json.NewDecoder(r.Body)
		enc := json.NewEncoder(w)
		err := dec.Decode(&context)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dbStore.AddContextTypeItem(context)

		section = dbStore.GetUserSectionBySectionsId(context.SectionsId)

		w.Header().Set("content-type", "application/json")
		enc.Encode(section)
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}
}

func DelContextItem(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(models.DatabaseStore); ok {
		context := new(models.ContextItem)
		var section models.Section

		dec := json.NewDecoder(r.Body)
		enc := json.NewEncoder(w)
		err := dec.Decode(&context)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dbStore.DelContextTypeItem(context)

		section = dbStore.GetUserSectionBySectionsId(context.SectionsId)

		w.Header().Set("content-type", "application/json")
		enc.Encode(section)
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}
}
