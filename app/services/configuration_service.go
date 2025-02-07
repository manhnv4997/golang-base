package services

import (
	"demo/app/utils"
	"fmt"
	"log"
)

type ConfigurationService struct{}

var (
	shopifyAPIKey      = utils.GetEnv("SHOPIFY_API_KEY", "")
	shopifyApiSecret   = utils.GetEnv("SHOPIFY_API_SECRET", "")
	shopifyRedirectURI = utils.GetEnv("SHOPIFY_REDIRECT_URI", "")
	shopifyScopes      = utils.GetEnv("SHOPIFY_SCOPES", "")
	client             = utils.NewClient()
)

func NewConfigurationService() *ConfigurationService {
	return &ConfigurationService{}
}

func (configurationService *ConfigurationService) Handle() (interface{}, error) {
	log.Println("OK....!!!")

	return "Xử lý thành công", nil
}

func (configurationService *ConfigurationService) GetAccessToken(shop string, code string) (string, error) {
	response, err := client.Post(
		fmt.Sprintf("https://%s/admin/oauth/access_token?client_id=%s&client_secret=%s&code=%s", shop, shopifyAPIKey, shopifyApiSecret, code),
		map[string]string{
			"client_id":     shopifyAPIKey,
			"client_secret": shopifyApiSecret,
			"code":          code,
		},
	)

	if err != nil {
		return "", err
	}

	// Trả về toàn bộ response.Body() dưới dạng chuỗi JSON
	return string(response.Body()), nil
}
