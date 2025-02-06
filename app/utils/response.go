package utils

import (
	"encoding/json"
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(response http.ResponseWriter, code int, message string) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(MessageResponse{Message: message})
	http.Error(response, message, code)
}

func SuccessResponse(response http.ResponseWriter, code int, data interface{}) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	encodeErr := json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "✅ Thành công",
		"result":  data,
	})

	if encodeErr != nil {
		return encodeErr
	}

	return nil
}
