package controllers

import (
	"demo/app/services"
	"encoding/json"
	"log"
	"net/http"
)

type UserController struct {
	Controllers
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (userController *UserController) GetAllUsers(response http.ResponseWriter, request *http.Request) {
	users, err := userController.UserService.GetAllUsers()
	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(response).Encode(map[string]interface{}{
		"message": "✅ Thành công",
		"result":  users,
	})

	if encodeErr != nil {
		http.Error(response, "Lỗi encode json:", http.StatusInternalServerError)
		log.Println("Lỗi encode json", encodeErr)
		return
	}
}
