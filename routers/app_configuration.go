package routers

import (
	"demo/app/controllers"
	"demo/app/services"
	"demo/app/utils"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupAppConfigurationRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	configurationService := services.NewConfigurationService()
	configurationController := controllers.NewConfigurationController(configurationService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc(utils.RoutePath("", ""), configurationController.Handle).Methods("GET")
	router.HandleFunc("/auth", configurationController.AuthHandler).Methods("GET")
	router.HandleFunc("/auth/callback", configurationController.CallbackHandler).Methods("GET")
	router.HandleFunc("/shop", configurationController.GetShopInfo).Methods("GET")

	return router
}
