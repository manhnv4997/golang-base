package routers

import (
	"demo/app/controllers"
	"demo/app/services"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupMenuRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	menuService := services.NewMenuService()
	menuController := controllers.NewMenuController(menuService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc("/menu", menuController.Store).Methods("POST")
	router.HandleFunc("/menu", menuController.Update).Methods("PUT")
	router.HandleFunc("/menu", menuController.Delete).Methods("DELETE")

	return router
}
