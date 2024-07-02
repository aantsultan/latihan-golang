package main

import (
	_ "github.com/lib/pq"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializerServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
