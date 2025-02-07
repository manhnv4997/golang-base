package services

import (
	"demo/app/utils"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type GraphQLService struct{}

func NewGraphQLService() *GraphQLService {
	return &GraphQLService{}
}

func (graphqlService *GraphQLService) GetProducts(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	shop := request.URL.Query().Get("shop")
	accessToken := request.URL.Query().Get("access_token")
	resp, err := utils.NewClient().Get(response, accessToken, fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/products.json", shop, utils.GetEnv("SHOPIFY_SCOPES", "2025-01")))

	return resp, err
}
