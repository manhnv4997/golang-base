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
	accessToken := "shpua_1b82388673bd287a2a085216c463880a"
	resp, err := utils.NewClient().Get(response, accessToken, fmt.Sprintf("https://%s/admin/api/%s/shop.json", "shop", "2025-01"))

	return resp, err
}
