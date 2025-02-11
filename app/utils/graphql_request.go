package utils

// Cấu trúc GraphQL request
type GraphQLRequest struct {
	Query string `json:"query"`
}

func GraphQLEndpoint(shop string) string {
	return "https://" + shop + ".myshopify.com/admin/api/" + GetEnv("SHOPIFY_DATE", "") + "/graphql.json"
}
