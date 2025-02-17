package services

import (
	"demo/app/utils"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ShopService struct{}

func NewShopService() *ShopService {
	return &ShopService{}
}

func (shopService *ShopService) GetDetail(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	// Lấy giá trị từ JSON
	resp, err := utils.NewClient().Get(response, fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/shop.json", "shop", "2025-01"))

	return resp, err
}
