package services

import (
	"demo/app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ThemeService struct{}

func NewThemeService() *ThemeService {
	return &ThemeService{}
}

func (themeService *ThemeService) Store(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	// Dữ liệu yêu cầu GraphQL
	query := `
		mutation themeCreate($source: URL!, $name: String!) {
			themeCreate(source: $source, name: $name) {
				theme {
					id
					name
					role
				}
				userErrors {
					field
					message
				}
			}
		}
	`

	// Biến cho GraphQL
	variables := map[string]interface{}{
		"source": bodyDataRequest["source"],
		"name":   bodyDataRequest["name"],
	}

	// Gửi yêu cầu POST
	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		map[string]interface{}{
			"query":     query,
			"variables": variables,
		},
	)

	// Kiểm tra lỗi
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	// Kiểm tra phản hồi
	if resp.IsError() {
		log.Fatalf("Error response from Shopify API: %s", resp.String())
	}

	// In ra kết quả thành công
	fmt.Println("Response from Shopify API:", resp.String())

	return resp, err
}

func (themeService *ThemeService) Update(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)
	if err != nil {
		return nil, err
	}

	query := `
    mutation themeUpdate($id: ID!, $input: OnlineStoreThemeInput!) {
        themeUpdate(id: $id, input: $input) {
            theme {
                id
                name
            }
            userErrors {
                field
                message
            }
        }
    }
    `

	// Tạo input object
	input := map[string]interface{}{
		"name": bodyDataRequest["name"],
	}

	// Tạo variables với ID và input
	variables := map[string]interface{}{
		"id":    fmt.Sprintf("gid://shopify/OnlineStoreTheme/%s", bodyDataRequest["id"]),
		"input": input,
	}

	// Gửi yêu cầu POST
	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		map[string]interface{}{
			"query":     query,
			"variables": variables,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		// Xử lý lỗi chi tiết như ví dụ trên
		return resp, fmt.Errorf("Error updating theme: %s", resp.String())
	}

	return resp, nil
}

func (themeService *ThemeService) Delete(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := `
    mutation themeDelete($id: ID!) {
        themeDelete(id: $id) {
			deletedThemeId
            userErrors {
                field
                message
            }
        }
    }
    `

	variables := map[string]interface{}{
		"id": fmt.Sprintf("gid://shopify/OnlineStoreTheme/%s", bodyDataRequest["id"]),
	}

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		map[string]interface{}{
			"query":     query,
			"variables": variables,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return resp, fmt.Errorf("Error updating theme: %s", resp.String())
	}

	return resp, nil
}

func (themeService *ThemeService) Publish(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := `
		mutation themePublish($id: ID!) {
			themePublish(id: $id) {
				theme {
					id
					name
				}
				userErrors {
					field
					message
				}
			}
		}
	`

	variables := map[string]interface{}{
		"id": fmt.Sprintf("gid://shopify/OnlineStoreTheme/%s", bodyDataRequest["id"]),
	}

	// Gửi yêu cầu POST
	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		map[string]interface{}{
			"query":     query,
			"variables": variables,
		},
	)

	// Kiểm tra lỗi
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// In ra kết quả thành công
	fmt.Println("Response from Shopify API:", resp.String())

	return resp, err
}
