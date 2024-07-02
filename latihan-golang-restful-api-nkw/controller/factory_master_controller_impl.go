package controller

import (
	"github.com/julienschmidt/httprouter"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/service"
	"net/http"
)

type FactoryMasterControllerImpl struct {
	Service service.FactoryMasterService
}

func NewFactoryMasterController(service service.FactoryMasterService) *FactoryMasterControllerImpl {
	return &FactoryMasterControllerImpl{
		Service: service,
	}
}

func (controller FactoryMasterControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	webResponse := controller.Service.FindAll(request.Context())
	helper.WriteToResponseBody(writer, webResponse)
}
