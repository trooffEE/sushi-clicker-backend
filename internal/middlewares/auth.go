package middlewares

import (
	"errors"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isUnprotectedRoute := lib.StringStartsWith(r.URL.Path, "/api/auth")

		if isUnprotectedRoute {
			next.ServeHTTP(w, r)
			return
		}

		reqToken := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		token := reqToken[1] // Removed bearer part
		if token == "" {
			http.Error(w, errors.New("invalid access token").Error(), http.StatusInternalServerError)
			return
		}

		_, err := lib.ValidateJwtAccessToken(token)
		if err != nil {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
