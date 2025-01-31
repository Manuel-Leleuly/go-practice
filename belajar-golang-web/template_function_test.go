package belajargolangweb

import (
	"belajar-golang-web/helpers"
	"net/http"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My name is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Manuel",
	})
}

func TestTemplateFunction(t *testing.T) {
	helpers.RunTemplateTest(TemplateFunction)
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Manuel",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	helpers.RunTemplateTest(TemplateFunctionGlobal)
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION").Funcs(map[string]any{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Manuel",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	helpers.RunTemplateTest(TemplateFunctionCreateGlobal)
}

func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION").Funcs(map[string]any{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Manuel",
	})
}

func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	helpers.RunTemplateTest(TemplateFunctionCreateGlobalPipeline)
}
