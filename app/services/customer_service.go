package services

import (
	"demo/app/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

type CustomerService struct {
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (customerService *CustomerService) List(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/customers.json", utils.GetEnv("SHOP_NAME", ""), utils.GetEnv("SHOPIFY_DATE", "2025-01")),
		)

	return resp, err
}

func (customerService *CustomerService) Detail(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	customerId := request.URL.Query().Get("customer-id")

	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/customers/%s.json", utils.GetEnv("SHOP_NAME", ""), utils.GetEnv("SHOPIFY_DATE", "2025-01"), customerId),
		)

	return resp, err
}

func (customerService *CustomerService) CustomerOrders(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	customerId := request.URL.Query().Get("customer-id")

	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/customers/%s/orders.json", utils.GetEnv("SHOP_NAME", ""), utils.GetEnv("SHOPIFY_DATE", "2025-01"), customerId),
		)

	return resp, err
}

func (customerService *CustomerService) Count(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	resp, err := utils.NewClient().
		Get(response,
			fmt.Sprintf("https://%s.myshopify.com/admin/api/%s/customers/count.json", utils.GetEnv("SHOP_NAME", ""), utils.GetEnv("SHOPIFY_DATE", "2025-01")),
		)

	return resp, err
}

type QueryParam struct {
	Key   string
	Value string
}

func (customerService *CustomerService) Search(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	queryParams := request.URL.Query()
	var paramArray []QueryParam

	for key, values := range queryParams {
		for _, value := range values {
			paramArray = append(paramArray, QueryParam{Key: key, Value: value})
		}
	}

	var paramString string
	var builder strings.Builder
	for _, param := range paramArray {
		builder.WriteString(param.Key)
		builder.WriteString(":")
		builder.WriteString(param.Value)
		builder.WriteString(",")
	}

	paramString = strings.TrimSuffix(builder.String(), ",")

	fmt.Println(paramString, "paramString")

	query := fmt.Sprintf(`{
		customers(first: 1, query: "%s"){
			edges {
				node {
					id
					firstName
					lastName
					email
				}
			}
		}
	}`, paramString,
	)

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (customerService *CustomerService) Store(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDatRequest, err := utils.BodyDataRequest(response, request)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`
		mutation {
			customerCreate(input: {
			firstName: "%s",
			lastName: "%s",
			email: "%s"
			})
			{
				customer{
					id
					firstName
					lastName
					email
				}
				userErrors{
					field
					message
				}
			}
		}
	`, bodyDatRequest["first_name"], bodyDatRequest["last_name"], bodyDatRequest["email"])

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (customerService *CustomerService) Update(response http.ResponseWriter, request *http.Request) (*resty.Response, error) {
	bodyDatRequest, err := utils.BodyDataRequest(response, request)

	log.Print(bodyDatRequest, "bodyDatRequest")

	if err != nil {
		return nil, err
	}

	dataArray := make(map[string]string)
	var inputString strings.Builder
	var fieldString strings.Builder
	for key, value := range bodyDatRequest {
		dataArray[key] = value.(string)

		inputString.WriteString(key)
		inputString.WriteString(":")
		inputString.WriteString(value.(string))
		inputString.WriteString(",")

		fieldString.WriteString(key)
		fieldString.WriteString(" ")
	}

	query := fmt.Sprintf(`
	mutation {
		customerUpdate(input: {
			%s
		}) 
		{
			customer {
				id
				firstName
				lastName
				email
			}
			userErrors {
				field
				message
			}
		}
	}`, inputString.String())

	log.Print(query, "query")

	resp, err := utils.NewClient().Post(
		utils.GraphQLEndpoint(utils.GetEnv("SHOP_NAME", "")),
		utils.GraphQLRequest{Query: query},
	)

	if err != nil {
		return nil, err
	}

	return resp, err
}
