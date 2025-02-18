package routers

import (
	"demo/app/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Setup routes
	router = SetupAppConfigurationRoutes(router)
	router = SetupShopRoutes(router)
	router = SetupUserRoutes(router)
	router = SetupProductRoutes(router)
	router = SetupCustomerRoutes(router)
	router = SetupPageRoutes(router)
	router = SetupMenuRoutes(router)
	router = SetupThemeRoutes(router)

	// Views
	router = InitViewRoutes(router)

	// In ra danh sách các routes
	utils.PrintRoutes(router)

	// Xử lý 404
	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("📌 Request: %s %s | Status: %d", request.Method, request.URL.Path, http.StatusNotFound)
		http.Error(response, "❌ Route không tồn tại", http.StatusNotFound)
	})

	return router
}
