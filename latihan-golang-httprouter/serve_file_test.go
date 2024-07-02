package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	expectedResult := "Hello World"
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(dir))

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/files/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, expectedResult, string(body))
}
