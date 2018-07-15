package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/roiperelman/client-site-server/models"
	"net/http"
	"strings"
)

type Claims struct{
	Id float64
}

const Secret = "secret string"

// AuthenticateDBUserMiddleware should continue where AuthenticateClaimsMiddleware took off
// should add User of type User context
func AuthenticateDBUserMiddleware(next http.Handler) http.Handler {
	return AuthenticateClaimsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if dbStore, ok := r.Context().Value("DBStore").(models.DatabaseStore); ok {
			var user models.User

			if id, ok := r.Context().Value("UserId").(int); ok {
				user.Id = id
				// check if user exists in db
				emailUser := dbStore.GetUserById(user.Id)
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
				http.Error(w, "authorization failed", http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), "User", user)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "internal DB problem", http.StatusInternalServerError)
			return
		}
	}))
}

// AuthenticateClaimsMiddleware should parse jwt, get claims and adds a UserId of type int context
func AuthenticateClaimsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			return []byte(Secret), nil
		})
		if err != nil {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}
		if id, ok := token.Claims.(jwt.MapClaims)["id"].(float64); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "UserId", int(id))
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "user doesnt exists", http.StatusUnauthorized)
			return
		}
	})
}