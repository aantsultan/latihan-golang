package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	expectedResult := "hello httprouter"
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, expectedResult)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, expectedResult, string(body))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		product := "Product " + id
		fmt.Fprint(writer, product)
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, "Product 1", string(body))
}
