package controllers

import (
	"demo/app/services"
	"demo/app/utils"
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

	encodeErr := utils.SuccessResponse(response, http.StatusOK, users)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}
