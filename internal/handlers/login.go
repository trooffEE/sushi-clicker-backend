package handlers

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.Write([]byte("Login"))
}
