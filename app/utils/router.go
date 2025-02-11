package utils

import (
	"fmt"

	"github.com/gorilla/mux"
)

func PrintRoutes(router *mux.Router) {
	fmt.Println("=== Danh sách Routes ===")
	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("Route:", pathTemplate)
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", methods)
		}
		fmt.Println("---")
		return nil
	})
	if err != nil {
		fmt.Println("Lỗi khi duyệt routes:", err)
	}
	fmt.Println("=======================")
}
