package latihan_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestFileServerEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	// StripPrefix digunakan untuk menghapus url path 'static' dalam /resources/static menjadi /resources
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{Addr: "localhost:8080",
		Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	// StripPrefix digunakan untuk menghapus url path 'static' dalam /resources/static menjadi /resources
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{Addr: "localhost:8080",
		Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
