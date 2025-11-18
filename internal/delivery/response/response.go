package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorStr string `json:"error"`
}

func SendErrorResponse(textError string, headerStatus int, w http.ResponseWriter) {
	response := ErrorResponse{ErrorStr: textError}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(headerStatus)
	json.NewEncoder(w).Encode(response)
}

func SendOKResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
