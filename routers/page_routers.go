package routers

import (
	"demo/app/controllers"
	"demo/app/services"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupPageRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	pageService := services.NewPageService()
	pageController := controllers.NewPageController(pageService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc("/page", pageController.Store).Methods("POST")
	router.HandleFunc("/page", pageController.Update).Methods("PUT")
	router.HandleFunc("/page", pageController.Delete).Methods("DELETE")

	return router
}
