package helper

import (
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
