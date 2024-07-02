package latihan_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/name.gohtml"))
	tmpl.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Struct title",
		Name:  "elaina",
		Address: Address{
			Street: "bekasi",
		},
	})
}

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/name.gohtml"))
	tmpl.ExecuteTemplate(writer, "name.gohtml",
		map[string]interface{}{
			"Title": "Template Title",
			"Name":  "aant",
			"Address": map[string]interface{}{
				"Street": "Bekasi",
			},
		})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
