package services

import (
	"demo/app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type PageService struct{}

func NewPageService() *PageService {
	return &PageService{}
}

func (pageService *PageService) Store(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	mutation {
		pageCreate(page: {
			title: "%s",
			body: "%s",
			handle: "%s"
		})
		{
			page{
				id
				title
				body
				handle
			}	
			userErrors{
				field
				message
			}
		}
	}
	`, bodyDataRequest["title"], bodyDataRequest["body_html"], bodyDataRequest["slug"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	log.Print(resp, "resp")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (pageService *PageService) Update(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	mutation {
		pageUpdate(
			id: "gid://shopify/Page/%s",
			page: {
				title: "%s",
				handle: "%s",
			}
		)
		{
			page{
				id
				title
				handle
			}
			userErrors{
				code
				field
				message
			}
		}
	}
	`, bodyDataRequest["id"], bodyDataRequest["title"], bodyDataRequest["slug"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (pageService *PageService) Delete(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	mutation {
		pageDelete(
			id: "gid://shopify/Page/%s",
		)
		{
			deletedPageId
			userErrors{
				field
				message
			}
		}
	}
	`, bodyDataRequest["id"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}
