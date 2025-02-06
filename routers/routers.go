package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//    // cách tạo instance
// type Routers struct{}

// func SetupRoutes() *mux.Router {
// 	router := mux.NewRouter()

// 	r := Routers{}   // cách tạo instance

// 	// Setup routes
// 	router = r.SetupUserRoutes(router)   // cách tạo instance

// 	// Xử lý 404
// 	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
// 		http.Error(response, "❌ Route không tồn tại", http.StatusNotFound)
// 	})

// 	return router
// }

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Setup routes
	router = SetupAppConfigurationRoutes(router)
	router = SetupUserRoutes(router)

	// Xử lý 404
	router.NotFoundHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("📌 Request: %s %s | Status: %d", request.Method, request.URL.Path, http.StatusNotFound)
		http.Error(response, "❌ Route không tồn tại", http.StatusNotFound)
	})

	return router
}
