package handlers

import (
	"encoding/json"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
)

func MultipleSectionsUser(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(*models.DBStore); ok {
		var isMulti bool

		dec := json.NewDecoder(r.Body)
		enc := json.NewEncoder(w)
		err := dec.Decode(&isMulti)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if id, ok := r.Context().Value("UserId").(int); ok {
			dbStore.UpdateIsMultipleSectionFeature(id, isMulti)
		} else {
			http.Error(w, "authorize user failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		enc.Encode(isMulti)
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}

}

func AddSectionUser(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(*models.DBStore); ok {
		var section models.Section

		dec := json.NewDecoder(r.Body)
		enc := json.NewEncoder(w)
		err := dec.Decode(&section)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if id, ok := r.Context().Value("UserId").(int); ok {
			sectionsId := dbStore.AddSection(id, section)
			section = dbStore.GetUserSectionBySectionsId(sectionsId)
		} else {
			http.Error(w, "authorize user failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		enc.Encode(section)
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}
}

func DelSectionUser(w http.ResponseWriter, r *http.Request) {
	if dbStore, ok := r.Context().Value("DBStore").(*models.DBStore); ok {
		var section models.Section

		dec := json.NewDecoder(r.Body)
		enc := json.NewEncoder(w)
		err := dec.Decode(&section)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if user, ok := r.Context().Value("User").(models.User); ok {
			if section.SectionId != user.DefaultSection {
				dbStore.DelSection(user.Id, section)
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
	} else {
		http.Error(w, "db connection failed", http.StatusInternalServerError)
		return
	}

}
