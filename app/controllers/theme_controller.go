package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type ThemeController struct {
	Controllers
	ThemeService *services.ThemeService
}

func NewThemeController(themeService *services.ThemeService) *ThemeController {
	return &ThemeController{ThemeService: themeService}
}

func (themeController *ThemeController) Store(response http.ResponseWriter, request *http.Request) {
	shop, err := services.NewThemeService().Store(response, request)

	if err != nil {
		log.Println(err, "err")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}
	log.Print(shop, "shop")

	bodyJson := utils.Decode(string(shop.Body()))
	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}

func (themeController *ThemeController) Update(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewThemeService().Update(response, request)

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

func (themeController *ThemeController) Delete(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewThemeService().Delete(response, request)

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

func (themeController *ThemeController) Publish(response http.ResponseWriter, request *http.Request) {
	result, err := services.NewThemeService().Publish(response, request)

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

// func (themeController *ThemeController) UpdateTheme(response http.ResponseWriter, request *http.Request) {
// 	result, err := services.NewThemeService().UpdateTheme(response, request)

// 	if err != nil {
// 		log.Println(err, "err")
// 		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
// 		return
// 	}

// 	bodyJson := utils.Decode(string(result.Body()))
// 	encodeErr := utils.SuccessResponse(response, http.StatusOK, bodyJson)

// 	if encodeErr != nil {
// 		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
// 		return
// 	}
// }
