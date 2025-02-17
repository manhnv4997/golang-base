package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type PageController struct {
	Controllers
	PageService *services.PageService
}

func NewPageController(pageService *services.PageService) *PageController {
	return &PageController{PageService: pageService}
}

func (pageController *PageController) Store(response http.ResponseWriter, request *http.Request) {
	shop, err := services.NewPageService().Store(response, request)

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

func (pageController *PageController) Update(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewPageService().Update(response, request)

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

func (pageController *PageController) Delete(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewPageService().Delete(response, request)

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
