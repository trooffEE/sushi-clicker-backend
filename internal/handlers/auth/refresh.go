package authHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/trooffEE/sushi-clicker-backend/internal/http"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"net/http"
)

var (
	RefreshError = errors.New("Refresh Error")
)

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie(httpServer.RefreshTokenName)
	if err != nil {
		http.Error(w, RefreshError.Error(), http.StatusUnauthorized)
		return
	}

	refreshToken, err := lib.ValidateJwtRefreshToken(tokenCookie.Value)
	if errors.Is(err, lib.InvalidTokenError) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "JWT claims missmatch", http.StatusInternalServerError)
	}
	email, sugar := claims["email"], claims["sugar"]

	accessToken, err := lib.GenerateJwtAccessToken(email.(string), sugar.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ResponseAccessToken{AccessToken: accessToken})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		fmt.Println(err)
	}
}
