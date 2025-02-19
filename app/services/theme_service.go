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

type ThemePublishResponse struct {
	Data struct {
		ThemePublish struct {
			Theme struct {
				ID string `json:"id"`
			} `json:"theme"`
			UserErrors []struct {
				Field   []string `json:"field"`
				Message string   `json:"message"`
			} `json:"userErrors"`
		} `json:"themePublish"`
	} `json:"data"`
}

func (themeService *ThemeService) Publish(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}
	mutation := `
    mutation themePublish($id: ID!) {
        themePublish(id: $id) {
            theme {
                id
            }
            userErrors {
                field
                message
            }
        }
    }
    `

	variables := map[string]interface{}{
		"id": fmt.Sprintf("gid://shopify/Theme/%s", bodyDataRequest["id"]),
	}

	payload := map[string]interface{}{
		"query":     mutation,
		"variables": variables,
	}

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		payload,
	)

	// Xử lý lỗi
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	// In kết quả phản hồi
	fmt.Println("Response Status:", resp.Status())
	fmt.Println("Response Body:", resp.String())

	return resp, err
}

func (themeService *ThemeService) UpdateTheme(response http.ResponseWriter, request http.Request) {
	// tạo theme

	// đặt theme vừa tạo làm mặc định

	// nhúng html
}
