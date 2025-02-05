package routers

import (
	"demo/app/controllers"
	"demo/app/middleware"
	"demo/app/repositories"
	"demo/app/services"
	"demo/app/utils"
	"demo/database/mysql"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func SetupUserRoutes(router *mux.Router) *mux.Router {
	// Khởi tạo Repository, Service, Controller
	userRepository := repositories.NewUserRepository(mysql.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	// Thêm middleware LoggerMiddleware
	router.Use(middleware.LoggerMiddleware) // Ví dụ sử dụng middleware cho toàn bộ route route
	// router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	// router.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userController.GetProfile))).Methods("GET")

	// Route
	router.HandleFunc(utils.RoutePath("user", ""), userController.GetAllUsers).Methods("GET")

	return router
}
