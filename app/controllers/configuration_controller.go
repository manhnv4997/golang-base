package controllers

import (
	"demo/app/services"
	"demo/app/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (configurationController *ConfigurationController) AuthHandler(response http.ResponseWriter, request *http.Request) {
	// https://7f39-14-191-163-181.ngrok-free.app/auth?shop=dungdinhnghe.myshopify.com
	// https://dungdinhnghe.myshopify.com/admin/oauth/authorize?client_id=7611fb8cd65cee6430d670f19bf58c22&scope=read_orders,write_products&redirect_uri=https://a59a-14-160-24-42.ngrok-free.app/auth/callback
	// https://dungdinhnghe.myshopify.com/admin/oauth/authorize?client_id=7611fb8cd65cee6430d670f19bf58c22&scope=read_orders,write_products&redirect_uri=https://a59a-14-160-24-42.ngrok-free.app/auth/callback&grant_options[]=per-user

	// https://dungdinhnghe.myshopify.com/admin/api/2025-01/products.json?ids=632910392921728736

	// Shopify OAuth URL
	authURL := fmt.Sprintf("https://%s.myshopify.com/admin/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s&state=%d", utils.GetEnv("SHOP_NAME", ""), shopifyAPIKey, shopifyScopes, shopifyRedirectURI, 123)

	log.Println(authURL, "authURL")

	http.Redirect(response, request, authURL, http.StatusFound)
}

func (configurationController *ConfigurationController) CallbackHandler(response http.ResponseWriter, request *http.Request) {
	shop := utils.GetEnv("SHOP_NAME", "")
	code := request.URL.Query().Get("code")
	hmac := request.URL.Query().Get("hmac")

	if code == "" {
		http.Error(response, "Invalid callback parameters", http.StatusBadRequest)
		return
	}

	// Verify HMAC (bảo mật)
	if !utils.ValidateHMAC(request.URL.Query(), shopifyApiSecret, hmac) {
		http.Error(response, "Invalid HMAC signature", http.StatusUnauthorized)
		return
	}

	// Step3: Exchange authorization code for access token
	body, err := services.NewConfigurationService().GetAccessToken(shop, code)

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
