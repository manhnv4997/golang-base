package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"log"
	"net/http"
)

type ShopController struct {
	Controllers
	ShopService *services.ShopService
}

func NewShopController(shopService *services.ShopService) *ShopController {
	return &ShopController{ShopService: shopService}
}

func (shopController *ShopController) GetDetail(response http.ResponseWriter, request *http.Request) {
	shop, err := services.NewShopService().GetDetail(response, request)

	if err != nil {
		log.Println(err, "errr")
		http.Error(response, "Lỗi lấy dữ liệu", http.StatusInternalServerError)
		return
	}

	encodeErr := utils.SuccessResponse(response, http.StatusOK, shop)

	if encodeErr != nil {
		utils.ErrorResponse(response, http.StatusInternalServerError, "Lỗi encode json")
		return
	}
}
