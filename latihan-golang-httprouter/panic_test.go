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

func TestPanicHandler(t *testing.T) {
	expectedResult := "Panic : ups"
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic : ", error)
	}
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("ups")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, expectedResult, string(body))
}
