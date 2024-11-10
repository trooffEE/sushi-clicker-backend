package httpServer

import (
	"github.com/trooffEE/sushi-clicker-backend/internal/config"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"net/http"
)

var RefreshTokenName = "X-Authorization-Refresh-Token"

func CookieInjectRefreshToken(w http.ResponseWriter, email string) {
	token, exp, err := lib.GenerateJwtRefreshToken(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     RefreshTokenName,
		Value:    token,
		Path:     "/",
		Expires:  exp,
		HttpOnly: true,
	}

	if !config.IsDevelopment {
		cookie.Secure = true
	}

	http.SetCookie(w, &cookie)
}
