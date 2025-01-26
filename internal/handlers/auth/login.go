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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var rUser user2.User
	err := json.NewDecoder(r.Body).Decode(&rUser)
	if err != nil {
		response.NewErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	usr, err := h.UserService.Login(rUser.Email, rUser.Password)
	if errors.Is(err, user.IncorrectCredentials) {
		response.NewErrorResponse(w, http.StatusUnauthorized, err)
		return
	}

	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	httpServer.CookieInjectRefreshToken(w, usr.Email, usr.Sugar)

	accessToken, err := lib.GenerateJwtAccessToken(usr.Email, usr.Sugar)
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	res, err := json.Marshal(ResponseAccessToken{AccessToken: accessToken})
	if err != nil {
		response.NewErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response.NewOkResponse(w, http.StatusOK, res)
}
