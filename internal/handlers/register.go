package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ConfirmPassword string `json:"confirmPassword"`
	Password        string `json:"password"`
	Email           string `json:"email"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if value, err := json.Marshal(user); err == nil {
		w.Write(value)
	}
}
