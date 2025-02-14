package main

import (
	"fmt"
	"net/http"

	"demo/app/utils"
	"demo/database/mysql"
	"demo/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Tải env
	utils.LoadEnv()

	// Kết nối database
	mysql.Connect()
	defer mysql.Close() //  ngắt kết nối database

	// Tải router
	router := routers.SetupRoutes()

	// Views
	router = routers.SetupViewRoutes(router)

	// Handle CORS
	handlerRouter := SetupCORS(router)

	// Khởi động server
	serverPort := ":" + utils.GetEnv("APP_PORT", "")
	fmt.Println("🚀 Server đang chạy trên cổng " + serverPort)
	http.ListenAndServe(serverPort, handlerRouter)
}

func SetupCORS(router *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://dungdinhnghe.myshopify.com"}, // Cho phép Shopify truy cập
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	return handler
}
