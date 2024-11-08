package auth

import (
	"encoding/json"
	"errors"
	"github.com/trooffEE/sushi-clicker-backend/internal/service/user"
	"net/http"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var usr user.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserService.Register(usr.Email, usr.Password)
	if errors.Is(err, user.IsAlreadyRegistered) {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	return
}
