package controllers

import (
	"demo/app/services"
	"fmt"
	"net/http"
)

type RoleController struct {
	Controllers
	UserService *services.UserService
}

func NewRoleController(userService *services.UserService) *RoleController {
	return &RoleController{UserService: userService}
}

func (roleController *RoleController) ShowHello(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello...")
}
