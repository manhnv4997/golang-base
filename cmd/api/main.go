package main

import (
	"fmt"
	"net/http"

	"demo/app/utils"
	"demo/database/mysql"
	"demo/routers"
)

func main() {
	// Tải env
	utils.LoadEnv()

	// Kết nối database
	mysql.Connect()
	defer mysql.Close() //  ngắt kết nối database

	// Tải router
	router := routers.SetupRoutes()

	// Khởi động server
	serverAddress := ":" + utils.GetEnv("APP_PORT", "")
	fmt.Println("🚀 Server đang chạy trên cổng " + serverAddress)
	http.ListenAndServe(serverAddress, router)
}
