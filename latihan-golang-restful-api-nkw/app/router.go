package app

import (
	"github.com/julienschmidt/httprouter"
	"latihan-golang-restful-api-nkw/controller"
)

func NewRouterFactoryMasters(factoryMasterController controller.FactoryMasterController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/factory-masters", factoryMasterController.FindAll)
	return router
}

func NewRouterWeavingKainMasters(controller controller.WeavingKainMasterController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/weaving-kain-masters", controller.FindAll)
	return router
}
