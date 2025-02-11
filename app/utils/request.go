package utils

import (
	"encoding/json"
	"net/http"
)

func BodyDataRequest(response http.ResponseWriter, request *http.Request) (map[string]interface{}, error) {
	var data map[string]interface{}

	// Giải mã JSON từ body vào struct
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(response, "Lỗi đọc JSON", http.StatusBadRequest)
		return nil, err
	}

	defer request.Body.Close()

	return data, nil
}
