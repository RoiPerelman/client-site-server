package handlers

import (
"net/http"
"encoding/json"
"github.com/roiperelman/client-site-server/models"
"strings"
"github.com/dgrijalva/jwt-go"
"fmt"
)

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

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Bearer" {
		http.Error(w, "authorization failed", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
		// Validate the alg is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			models.AddSection(int(claims["id"].(float64)), int(section))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	w.Header().Set("content-type", "application/json")
	enc.Encode(section)
}
