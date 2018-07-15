package middlewares

import (
	"net/http"
	"context"
	"github.com/roiperelman/client-site-server/models"
)

type DBStore struct {
	models.DatabaseStore
}
func (dbStore *DBStore) DBStoreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "DBStore", dbStore)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}