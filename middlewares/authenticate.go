package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
	"strings"
)

const secret = "secret string"

// todo split into 2 middlewares - one that only checks the jwt token claims and adds them to context "tokenClaims" and another that gets User from db and adds it to context under User.
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				user.Id = int(claims["id"].(float64))
				// check if user exists in db
				emailUser := models.GetUserById(user.Id)
				if emailUser != nil {
					user.IsAuthenticated = true
					user.Id = emailUser.Id
					user.IsMulti = emailUser.IsMulti
					user.DefaultSection = emailUser.DefaultSection
					user.Sections = emailUser.Sections
					user.Username = emailUser.Username
					user.Email = emailUser.Email
					user.JSCode = emailUser.JSCode
				} else {
					http.Error(w, "user doesnt exists", http.StatusUnauthorized)
					return
				}
			} else {
				http.Error(w, "user doesnt exists", http.StatusUnauthorized)
				return
			}
		}

		ctx := context.WithValue(r.Context(), "User", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
