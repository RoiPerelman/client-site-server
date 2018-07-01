package middlewares

import (
	"net/http"
	"context"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/roiperelman/client-site-server/models"
)

const secret = "secret string"

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
					user.Id = emailUser.Id
					user.IsMulti = emailUser.IsMulti
					user.DefaultSection = emailUser.DefaultSection
					user.Sections = emailUser.Sections
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				user.IsAuthenticated = false
				w.WriteHeader(http.StatusUnauthorized)
			}
		}

		ctx := context.WithValue(r.Context(), "User", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}