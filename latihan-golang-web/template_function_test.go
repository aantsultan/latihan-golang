package latihan_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "aant"}}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "aant",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	//tmpl := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	tmpl := template.New("FUNCTION")
	// create custom global function
	tmpl = tmpl.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	tmpl = template.Must(tmpl.Parse(`{{ upper .Name }}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "aant",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipeLines(writer http.ResponseWriter, request *http.Request) {
	//tmpl := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	tmpl := template.New("FUNCTION")
	// create custom global function
	tmpl = tmpl.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	tmpl = template.Must(tmpl.Parse(`{{ sayHello .Name | upper }}`))
	tmpl.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "aant",
	})
}

func TestTemplateFunctionPipeLines(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipeLines(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
