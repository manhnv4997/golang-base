package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type MenuController struct {
	Controllers
	MenuService *services.MenuService
}

func NewMenuController(menuService *services.MenuService) *MenuController {
	return &MenuController{MenuService: menuService}
}

func (menuController *MenuController) Store(response http.ResponseWriter, request *http.Request) {
	shop, err := services.NewMenuService().Store(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(shop.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (menuController *MenuController) Update(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewMenuService().Update(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(result.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (menuController *MenuController) Delete(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewMenuService().Delete(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	bodyJson := utils.Decode(string(result.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}
