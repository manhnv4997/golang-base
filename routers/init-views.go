package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// func (r *Routers) SetupUserRoutes(router *mux.Router) *mux.Router {   // cách tạo instance
func InitViewRoutes(router *mux.Router) *mux.Router {

	// Serve file tĩnh từ thư mục /view/public
	fs := http.FileServer(http.Dir("./view/public"))
	router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

	return router
}
