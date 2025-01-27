package authHandler

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	httpServer "github.com/trooffEE/sushi-clicker-backend/internal/http"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"github.com/trooffEE/sushi-clicker-backend/internal/response"
	"net/http"
)

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie(httpServer.RefreshTokenName)
	if err != nil {
		response.NewErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	refreshToken, err := lib.ValidateJwtRefreshToken(tokenCookie.Value)
	if errors.Is(err, lib.InvalidTokenError) {
		response.NewErrorResponse(w, http.StatusUnauthorized, err)
		return
	}
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		response.NewErrorResponse(w, http.StatusInternalServerError, errors.New("JWT claims missmatch"))
	}
	email, sugar := claims["email"], claims["sugar"]

	accessToken, err := lib.GenerateJwtAccessToken(email.(string), sugar.(string))
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.NewOkResponse(w, http.StatusOK, ResponseAccessToken{AccessToken: accessToken})
}
