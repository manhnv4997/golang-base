package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type ConfigurationController struct {
	Controllers
	ConfigurationService *services.ConfigurationService
}

var (
	shopifyAPIKey      = utils.GetEnv("SHOPIFY_API_KEY", "")
	shopifyApiSecret   = utils.GetEnv("SHOPIFY_API_SECRET", "")
	shopifyRedirectURI = utils.GetEnv("SHOPIFY_REDIRECT_URI", "")
	shopifyScopes      = utils.GetEnv("SHOPIFY_SCOPES", "")
)

func NewConfigurationController(configurationService *services.ConfigurationService) *ConfigurationController {
	return &ConfigurationController{ConfigurationService: configurationService}
}

func (configurationController *ConfigurationController) Handle(response http.ResponseWriter, request *http.Request) {
	result, err := configurationController.ConfigurationService.Handle()

	if err != nil {
		log.Println(err, "err")
		ErrorResponse(response, http.StatusInternalServerError, "Lỗi xử lý dữ liệu")
		return
	}

	SuccessResponse(response, http.StatusOK, result)
}

func (configurationController *ConfigurationController) AuthHandler(response http.ResponseWriter, request *http.Request) {
	shop := request.URL.Query().Get("shop")
	if shop == "" {
		// ErrorResponse(response, http.StatusBadRequest, "Missing shop parameter")

		// http.Error(response, "Missing shop parameter", http.StatusBadRequest)
		ErrorResponse(response, http.StatusBadRequest, "Missing shop parameter")
		return
	}

	// Shopify OAuth URL
	authURL := fmt.Sprintf("https://%s/admin/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s", shop, shopifyAPIKey, shopifyScopes, shopifyRedirectURI)

	http.Redirect(response, request, authURL, http.StatusFound)
}

func (configurationController *ConfigurationController) CallbackHandler(response http.ResponseWriter, request *http.Request) {
	shop := request.URL.Query().Get("shop")
	code := request.URL.Query().Get("code")
	hmac := request.URL.Query().Get("hmac")

	if shop == "" || code == "" {
		http.Error(response, "Invalid callback parameters", http.StatusBadRequest)
		return
	}

	// Verify HMAC (bảo mật)
	if !validateHMAC(request.URL.Query(), shopifyApiSecret, hmac) {
		http.Error(response, "Invalid HMAC signature", http.StatusUnauthorized)
		return
	}

	// Step3: Exchange authorization code for access token
	accessToken, err := getAccessToken(shop, code)

	if err != nil {
		http.Error(response, "Failed to get access token", http.StatusInternalServerError)
		return
	}

	// Return access token (trong thưc tế bạn nên lưu vào database)
	result := map[string]string{"shop": shop, "access_token": accessToken}
	// json.NewEncoder(response).Encode(result)

	SuccessResponse(response, http.StatusOK, result)
}

func validateHMAC(param map[string][]string, secret, receivedHMAC string) bool {
	// Logic xác thực HMAC (có thể dùng crypto/hmac)
	return true // Giả định HMAC hợp lệ
}

func getAccessToken(shop, code string) (string, error) {
	client := resty.New()        // Tạo HTTP client
	response, err := client.R(). // Tạo một request HTTP
					SetHeader("Content-Type", "application/json").
					SetBody(map[string]string{
			"client_id":     shopifyAPIKey,
			"client_secret": shopifyApiSecret,
			"code":          code,
		}).
		Post(fmt.Sprintf("https://%s/admin/oauth/access_token", shop)) // Gửi request GET

	if err != nil {
		return "", err
	}

	var result map[string]string
	if err := json.Unmarshal(response.Body(), result); err != nil {
		return "", err
	}

	return result["access_token"], nil
}

// Step 4: Use Access token to call shopify API
func (configurationController *ConfigurationController) GetShopInfo(response http.ResponseWriter, request *http.Request) {
	shop := request.URL.Query().Get("shop")
	accessToken := request.URL.Query().Get("access_token")

	if shop == "" || accessToken == "" {
		http.Error(response, "Missing shop or access_token", http.StatusBadRequest)
		return
	}

	client := resty.New()    // Tạo HTTP Client
	resp, err := client.R(). // Tạo một request HTTP
					SetAuthToken(accessToken).
					Get(fmt.Sprintf("https://%s/admin/api/2023-04/shop.json", "shop"))

	if err != nil {
		http.Error(response, "Failed to fetch shop info", http.StatusInternalServerError)
		return
	}

	response.Write(resp.Body())
}
