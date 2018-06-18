package handlers

import (
	"net/http"
	"github.com/roiperelman/client-site-server/models"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"encoding/json"
)

func AuthorizeUser(	w http.ResponseWriter, r *http.Request) {
	var user models.User

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
		user.Errors.Server = err.Error()
		user.IsAuthenticated = false
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user.Username = claims["username"].(string)
			user.Email = claims["email"].(string)

			// check if user exists in db
			emailUser := models.GetUserByEmail(user.Email)
			if emailUser != nil && emailUser.Username == user.Username {
				user.IsAuthenticated = true
				user.SectionId = emailUser.SectionId
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			user.IsAuthenticated = false
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	// Marshal the struct to []byte format
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// Write the output
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}