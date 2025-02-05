package middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("📌 Request: %s %s", request.Method, request.URL.Path)
		next.ServeHTTP(response, request)
	})
}
