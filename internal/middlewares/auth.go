package middlewares

import (
	"errors"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"github.com/trooffEE/sushi-clicker-backend/internal/response"
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
			response.NewErrorResponse(w, http.StatusBadRequest, errors.New("no access token"))
			return
		}

		_, err := lib.ValidateJwtAccessToken(token)
		if err != nil {
			response.NewErrorResponse(w, http.StatusUnauthorized, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
