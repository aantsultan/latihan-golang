package test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io"
	"latihan-golang-restful-api-nkw/app"
	"latihan-golang-restful-api-nkw/controller"
	"latihan-golang-restful-api-nkw/middleware"
	"latihan-golang-restful-api-nkw/repository"
	"latihan-golang-restful-api-nkw/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRouter(db *sql.DB) http.Handler {
	fmRepository := repository.NewFactoryMasterRepository()
	fmService := service.NewFactoryMasterService(fmRepository, db)
	fmController := controller.NewFactoryMasterController(fmService)

	router := app.NewRouter(fmController)
	return middleware.NewAuthMiddleware(router)
}

func TestGetListFactoryMaster(t *testing.T) {
	db := app.NewDB()
	router := SetupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/factory-masters", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

}
