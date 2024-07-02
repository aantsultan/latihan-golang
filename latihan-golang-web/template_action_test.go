package latihan_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/address.gohtml"))
	tmpl.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Name": "Aant",
		//"Address": map[string]interface{}{
		//	"Street": "Perum Lia",
		//	"City":   "Bekasi",
		//},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/range.gohtml"))
	tmpl.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{"Game", "Manga", "Anime"},
	})
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionComparator(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	tmpl.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"FinalValue": 50,
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/if.gohtml"))
	tmpl.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Struct title",
		//Name:  "elaina",
	})
}
