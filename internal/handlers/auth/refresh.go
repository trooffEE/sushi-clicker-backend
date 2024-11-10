package authHandler

import (
	"errors"
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
	}

	_, err := lib.ValidateJwtRefreshToken(tokenCookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//accessToken, err := lib.GenerateJwtAccessToken(usr.Email, usr.Sugar)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//response, err := json.Marshal(Response{accessToken: accessToken})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	//if _, err := w.Write(response); err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
}
