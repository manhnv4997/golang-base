package main

import (
	"fmt"
	"net/http"

	"demo/app/utils"
	"demo/database/mysql"
	"demo/routers"
)

func main() {
	mysql.Connect("root", "", "127.0.0.1", "3306", "shopone")
	defer mysql.Close()

	router := routers.SetupRoutes()

	// Khởi động server
	serverAddress := ":" + utils.GetEnv("APP_PORT", "")
	fmt.Println("🚀 Server đang chạy trên cổng " + serverAddress)
	http.ListenAndServe(serverAddress, router)
}
