package response

import (
	"net/http"
)

type ErrorResponse struct {
	Success bool   "json:success"
	Reason  string "json:reason"
	Code    int64  "json:reason"
}

func NewErrorResponse(w http.ResponseWriter, code int, err error) {
	response := marshall(ErrorResponse{
		Success: false,
		Reason:  err.Error(),
		Code:    int64(code),
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
