package response

import (
	"net/http"
)

type OkResponse struct {
	Success bool        "json:success"
	Payload interface{} "json:payload"
	Code    int64       "json:reason"
}

func NewOkResponse(w http.ResponseWriter, code int, payload interface{}) {
	response := marshall(OkResponse{
		Code:    int64(code),
		Success: true,
		Payload: payload,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
