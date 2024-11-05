package handlers

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login"))
}
