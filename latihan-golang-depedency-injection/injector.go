//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"latihan-golang-depedency-injection/app"
	"latihan-golang-depedency-injection/controller"
	"latihan-golang-depedency-injection/middleware"
	"latihan-golang-depedency-injection/repository"
	"latihan-golang-depedency-injection/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer(option ...validator.Option) *http.Server {
	wire.Build(
		app.NewDB,
		//wire.Bind(new([]validator.Option), new(*(interface{})(nil))),
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
