package routers

import (
	"demo/app/controllers"
	"demo/app/services"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupThemeRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	themeService := services.NewThemeService()
	themeController := controllers.NewThemeController(themeService)

	// Thêm middleware LoggerMiddleware
	// router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc("/theme", themeController.Store).Methods("POST")
	router.HandleFunc("/theme", themeController.Update).Methods("PUT")
	router.HandleFunc("/theme", themeController.Delete).Methods("DELETE")
	router.HandleFunc("/theme/publish", themeController.Publish).Methods("POST")
	// router.HandleFunc("/update-theme", themeController.UpdateTheme).Methods("POST")

	return router
}
