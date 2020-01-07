package api

import (
	"context"
	"net/http"
	"rest_api_template/internal/app/models"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Api-Auth-Token")

		if uid := models.UserGetByToken(token); uid != 0 {
			ctx := context.WithValue(r.Context(), "uid", uid)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	})
}
