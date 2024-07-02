package latihan_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World template")
}

func TestSimpleHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	tmpl.ExecuteTemplate(writer, "simple.gohtml", "Hello World template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//parse, err := templates.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	parse := template.Must(template.New("SIMPLE").Parse(templateText))

	err := parse.ExecuteTemplate(writer, "SIMPLE", "Hello html templates")
	if err != nil {
		panic(err)
	}
}
