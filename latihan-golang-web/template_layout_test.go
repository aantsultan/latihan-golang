package latihan_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml"))
	tmpl.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Name": "Aant",
		//"Address": map[string]interface{}{
		//	"Street": "Perum Lia",
		//	"City":   "Bekasi",
		//},
	})
}
