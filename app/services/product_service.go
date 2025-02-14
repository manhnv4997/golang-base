package services

import (
	"demo/app/utils"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (productService *ProductService) List(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	shop := utils.GetEnv("SHOP_NAME", "")

	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/products.json", shop, utils.GetEnv("SHOPIFY_DATE", "2025-01")),
		)

	return resp, err
}

func (productService *ProductService) Detail(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	shop := utils.GetEnv("SHOP_NAME", "")
	id := request.URL.Query().Get("product-id")

	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/products/%s.json", shop, utils.GetEnv("SHOPIFY_DATE", "2025-01"), id),
		)

	return resp, err
}

func (productService *ProductService) CountProduct(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	shop := utils.GetEnv("SHOP_NAME", "")

	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/products/count.json", shop, utils.GetEnv("SHOPIFY_DATE", "2025-01")),
		)

	return resp, err
}

func (productService *ProductService) Update(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	// Đọc toàn bộ body request
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	// GraphQL mutation để cập nhật sản phẩm
	query := fmt.Sprintf(`
	mutation {
		productUpdate(input: { 
			id: "gid://shopify/Product/%s", 
			title: "%s" 
		}) {
			product {
				id
				title
			}
			userErrors {
				field
				message
			}
		}
	}`, bodyDataRequest["product_id"], bodyDataRequest["title"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(bodyDataRequest["shop"].(string)),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (productService *ProductService) Delete(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	productId := request.URL.Query().Get("product_id")
	shop := utils.GetEnv("SHOP_NAME", "")

	query := fmt.Sprintf(`
		mutation {
			productDelete(input: {id: %q}) {
				deletedProductId
				userErrors {
					field
					message
				}
			}
		}`, fmt.Sprintf("gid://shopify/Product/%s", productId))

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(shop),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}
