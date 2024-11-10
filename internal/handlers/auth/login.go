package authHandler

import (
	"encoding/json"
	"errors"
	"fmt"
	httpServer "github.com/trooffEE/sushi-clicker-backend/internal/http"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
	"net/http"
)

type Response struct {
	accessToken string "json:accessToken"
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var rUser user.User
	err := json.NewDecoder(r.Body).Decode(&rUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usr, err := h.UserService.Login(rUser.Email, rUser.Password)
	if errors.Is(err, user.IncorrectCredentials) {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpServer.CookieInjectRefreshToken(w, usr.Email)

	accessToken, err := lib.GenerateJwtAccessToken(usr.Email, usr.Sugar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(Response{accessToken: accessToken})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		fmt.Println(err)
	}
}
