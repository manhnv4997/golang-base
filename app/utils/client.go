package utils

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct{}

var (
	shopifyAPIKey      = GetEnv("SHOPIFY_API_KEY", "")
	shopifyApiSecret   = GetEnv("SHOPIFY_API_SECRET", "")
	shopifyRedirectURI = GetEnv("SHOPIFY_REDIRECT_URI", "")
	shopifyScopes      = GetEnv("SHOPIFY_SCOPES", "")
)

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Get(response http.ResponseWriter, accessToken string, url string) (*resty.Response, error) {
	clt := resty.New()    // Tạo HTTP Client
	resp, err := clt.R(). // Tạo một request HTTP
				SetAuthToken(accessToken).
				Get(url)

	return resp, err
}

func (client *Client) Post(url string, bodyData map[string]string) (*resty.Response, error) {
	clt := resty.New()    // Tạo HTTP client
	resp, err := clt.R(). // Tạo một request HTTP
				SetHeader("Content-Type", "application/json").
				SetBody(bodyData).
				Post(url) // Gửi request POST

	if err != nil {
		return resp, err
	}

	// Trả về toàn bộ response.Body() dưới dạng chuỗi JSON
	return resp, nil
}
