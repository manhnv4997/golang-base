package utils

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Client struct{}

var (
	accessToken = GetEnv("SHOPIFY_ACCESS_TOKEN", "")
)

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Get(response http.ResponseWriter, url string) (*resty.Response, error) {
	clt := resty.New()    // Tạo HTTP Client
	resp, err := clt.R(). // Tạo một request HTTP
				SetHeader("Content-Type", "application/json").
				SetHeader("X-Shopify-Access-Token", accessToken).
				Get(url)

	return resp, err
}

// func (client *Client) Post(url string, bodyData map[string]string) (*resty.Response, error) {
func (client *Client) Post(url string, bodyData interface{}) (*resty.Response, error) {
	clt := resty.New()    // Tạo HTTP client
	resp, err := clt.R(). // Tạo một request HTTP
				SetHeader("Content-Type", "application/json").
				SetHeader("X-Shopify-Access-Token", accessToken).
				SetBody(bodyData).
				Post(url) // Gửi request POST

	if err != nil {
		return resp, err
	}

	// Trả về toàn bộ response.Body() dưới dạng chuỗi JSON
	return resp, nil
}
