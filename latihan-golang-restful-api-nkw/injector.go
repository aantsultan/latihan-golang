//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"latihan-golang-restful-api-nkw/app"
	"latihan-golang-restful-api-nkw/controller"
	"latihan-golang-restful-api-nkw/middleware"
	"latihan-golang-restful-api-nkw/repository"
	"latihan-golang-restful-api-nkw/service"
	"net/http"
)

var weavingKainMasterSet = wire.NewSet(
	repository.NewWeavingKainMasterRepository,
	wire.Bind(new(repository.WeavingKainMasterRepository), new(*repository.WeavingKainMasterRepositoryImpl)),
	service.NewFactoryMasterService,
	wire.Bind(new(service.WeavingKainMasterService), new(*service.WeavingKainMasterServiceImpl)),
	controller.NewWeavingKainMasterController,
	wire.Bind(new(controller.WeavingKainMasterController), new(*controller.WeavingKainMasterControllerImpl)),
)

var factoryMasterSet = wire.NewSet(
	repository.NewFactoryMasterRepository,
	wire.Bind(new(repository.FactoryMasterRepository), new(*repository.FactoryMasterRepositoryImpl)),
	service.NewFactoryMasterService,
	wire.Bind(new(service.FactoryMasterService), new(*service.FactoryMasterServiceImpl)),
	controller.NewFactoryMasterController,
	wire.Bind(new(controller.FactoryMasterController), new(*controller.FactoryMasterControllerImpl)),
)

func InitializerServer() *http.Server {
	wire.Build(
		app.NewDB,
		factoryMasterSet,
		app.NewRouterFactoryMasters,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
