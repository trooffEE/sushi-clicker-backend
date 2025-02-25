package authHandler

import (
	"encoding/json"
	"errors"
	httpServer "github.com/trooffEE/sushi-clicker-backend/internal/http"
	"github.com/trooffEE/sushi-clicker-backend/internal/lib"
	user2 "github.com/trooffEE/sushi-clicker-backend/internal/models/user"
	"github.com/trooffEE/sushi-clicker-backend/internal/response"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var usr user2.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usrDb, err := h.UserService.Register(usr.Email, usr.Password)
	if errors.Is(err, user.IsAlreadyRegistered) {
		response.NewErrorResponse(w, http.StatusConflict, errors.New("user is already registered"))
		return
	}
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, errors.New("something went wrong"))
		return
	}

	httpServer.CookieInjectRefreshToken(w, usr.Email, usrDb.Sugar)

	accessToken, err := lib.GenerateJwtAccessToken(usr.Email, usrDb.Sugar)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.NewOkResponse(w, http.StatusCreated, ResponseAccessToken{AccessToken: accessToken})
	return
}
