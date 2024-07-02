package app

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"latihan-golang-restful-api/controller"
	"latihan-golang-restful-api/exception"
	"latihan-golang-restful-api/helper"
	"time"
)

func SetupTestDb() *sql.DB {
	db, err := sql.Open("mysql", "root:Sabeso76@tcp(localhost:3306)/latihan_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
