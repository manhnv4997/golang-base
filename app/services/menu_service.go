package services

import (
	"demo/app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type MenuService struct{}

func NewMenuService() *MenuService {
	return &MenuService{}
}

func (menuService *MenuService) Store(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	log.Print(bodyDataRequest, "bodyDataRequest")

	if err != nil {
		return nil, err
	}

	// GraphQL Mutation
	query := `mutation {
		menuCreate(
			title: "Main Menu",
			handle: "main-menu",
			items: [
				{
					title: "Home",
					url: "/",
					type: HTTP
				},
				{
					title: "Products",
					url: "/collections/all",
					type: HTTP
				}
			]
		) {
			menu {
				id
				title
				items {
					id
					title
					url
				}
			}
			userErrors {
				field
				message
			}
		}
	}`

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (menuService *MenuService) Update(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	log.Print(bodyDataRequest, "bodyDataRequest")

	query := fmt.Sprintf(`mutation {
		menuUpdate(
			id: "gid://shopify/Menu/%s",
			title: "Updated Main Menu 1",
			handle: "updated-main-menu",
			items: [
				{
					title: "Home",
					url: "/",
					type: HTTP
				},
				{
					title: "Shop",
					url: "/collections/all",
					type: HTTP
				},
				{
					title: "Contact",
					url: "/contact",
					type: HTTP
				}
			]
		) {
			menu {
				id
				title
				items {
					id
					title
					url
				}
			}
			userErrors {
				field
				message
			}
		}
	}`, bodyDataRequest["id"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (menuService *MenuService) Delete(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDataRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
	mutation {
		menuDelete(
			id: "gid://shopify/Menu/%s",
		)
		{
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
