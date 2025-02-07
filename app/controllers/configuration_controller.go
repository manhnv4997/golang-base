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
		// http.Error(response, "Missing shop parameter", http.StatusBadRequest)
		ErrorResponse(response, http.StatusBadRequest, "Missing shop parameter")
		return
	}

	// https://a59a-14-160-24-42.ngrok-free.app/auth?shop=dungdinhnghe.myshopify.com
	// https://dungdinhnghe.myshopify.com/admin/oauth/authorize?client_id=7611fb8cd65cee6430d670f19bf58c22&scope=read_orders,write_products&redirect_uri=https://a59a-14-160-24-42.ngrok-free.app/auth/callback
	// https://dungdinhnghe.myshopify.com/admin/oauth/authorize?client_id=7611fb8cd65cee6430d670f19bf58c22&scope=read_orders,write_products&redirect_uri=https://a59a-14-160-24-42.ngrok-free.app/auth/callback&grant_options[]=per-user

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
	if !utils.ValidateHMAC(request.URL.Query(), shopifyApiSecret, hmac) {
		http.Error(response, "Invalid HMAC signature", http.StatusUnauthorized)
		return
	}

	// Step3: Exchange authorization code for access token
	body, err := getAccessToken(shop, code)

	// In toàn bộ JSON nhận được
	log.Println("Response Body:", body)

	// Parse JSON từ chuỗi
	var bodyJson map[string]interface{}
	if err := json.Unmarshal([]byte(body), &bodyJson); err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	// Lấy giá trị từ JSON
	accessToken := bodyJson["access_token"].(string)
	scope := bodyJson["scope"].(string)

	if err != nil {
		ErrorResponse(response, http.StatusInternalServerError, "Failed to get access token")
		return
	}

	// Return access token (trong thưc tế bạn nên lưu vào database)
	result := map[string]string{"shop": shop, "access_token": accessToken, "scope": scope}

	SuccessResponse(response, http.StatusOK, result)
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
		Post(fmt.Sprintf("https://%s/admin/oauth/access_token?client_id=%s&client_secret=%s&code=%s", shop, shopifyAPIKey, shopifyApiSecret, code)) // Gửi request GET

	if err != nil {
		return "", err
	}

	// Trả về toàn bộ response.Body() dưới dạng chuỗi JSON
	return string(response.Body()), nil
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
