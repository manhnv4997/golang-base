package routers

import (
	"demo/app/controllers"
	"demo/app/services"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupProductRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	productService := services.NewProductService()
	productController := controllers.NewProductController(productService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc("/product", productController.List).Methods("GET")
	router.HandleFunc("/product/detail", productController.Detail).Methods("GET")
	router.HandleFunc("/product/count", productController.CountProduct).Methods("GET")
	router.HandleFunc("/product/update", productController.Update).Methods("PUT")
	router.HandleFunc("/product", productController.Delete).Methods("DELETE")
	// router.HandleFunc("/product/count", productController.CountProduct).Methods("GET")

	return router
}
