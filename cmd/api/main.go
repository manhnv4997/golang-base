package main

import (
	"fmt"
	"net/http"

	"demo/database/mysql"
	"demo/routers"
)

func main() {
	mysql.Connect("root", "", "127.0.0.1", "3306", "shopone")
	defer mysql.Close()

	router := routers.SetupRoutes()

	// Khởi động server
	fmt.Println("🚀 Server đang chạy trên cổng 8080")
	http.ListenAndServe(":8080", router)
}
