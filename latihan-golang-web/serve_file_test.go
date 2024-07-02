package latihan_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

//go:embed resources/index.html
var resourceIndex string

//go:embed resources/pages/sign-in.html
var resourceSignIn string

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceIndex)
	} else {
		fmt.Fprint(writer, resourceSignIn)
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		http.ServeFile(writer, request, "./resources/pages/sign-in.html")
	}
}
