package routers

import (
	"demo/app/controllers"
	"demo/app/services"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupCustomerRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	customerService := services.NewCustomerService()
	customerController := controllers.NewCustomerController(customerService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc("/customer", customerController.Store).Methods("POST")
	router.HandleFunc("/customer", customerController.List).Methods("GET")
	router.HandleFunc("/customer", customerController.Update).Methods("PUT")
	router.HandleFunc("/customer/detail", customerController.Detail).Methods("GET")
	router.HandleFunc("/customer/orders", customerController.CustomerOrders).Methods("GET")
	router.HandleFunc("/customer/count", customerController.Count).Methods("GET")
	router.HandleFunc("/customer/search", customerController.Search).Methods("GET")

	return router
}
