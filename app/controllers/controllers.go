package controllers

import (
	"demo/app/utils"
	"net/http"
)

type Controllers struct {
}

func SuccessResponse(response http.ResponseWriter, code int, data interface{}) {
	encodeErr := utils.SuccessResponse(response, code, data)

	if encodeErr != nil {
		ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
	}
}

func ErrorResponse(response http.ResponseWriter, code int, message string) {
	utils.ErrorResponse(response, code, message)
}
