package middlewares

import (
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isUnprotectedRoute := lib.StringStartsWith(r.URL.Path, "/api/user")

		if isUnprotectedRoute {
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("X-Authorization-Access-Token")
		if err != nil {
			//http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		_, err = lib.ValidateJwtAccessToken(token.Value)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		// Check protection
		next.ServeHTTP(w, r)
	})
}
