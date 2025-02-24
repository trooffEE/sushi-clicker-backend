package middlewares

import (
	"go.uber.org/zap"
	"net/http"
)

func HTTPHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		zap.L().Info("Serving request", zap.String("method", r.Method), zap.String("endpoint", r.RequestURI))
		next.ServeHTTP(w, r)
	})
}
