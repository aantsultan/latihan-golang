package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type FactoryMasterController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
