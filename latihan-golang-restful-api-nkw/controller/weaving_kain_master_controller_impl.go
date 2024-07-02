package controller

import (
	"github.com/julienschmidt/httprouter"
	"latihan-golang-restful-api-nkw/helper"
	"latihan-golang-restful-api-nkw/service"
	"net/http"
)

type WeavingKainMasterControllerImpl struct {
	Service service.WeavingKainMasterService
}

func NewWeavingKainMasterController(service service.WeavingKainMasterService) *WeavingKainMasterControllerImpl {
	return &WeavingKainMasterControllerImpl{Service: service}
}

func (controller WeavingKainMasterControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	webResponse := controller.Service.FindAll(request.Context())
	helper.WriteToResponseBody(writer, webResponse)
}
