package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"text/template"
)

func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml",
	))

	t.ExecuteTemplate(writer, "layout", map[string]any{
		"Name":  "Manuel",
		"Title": "Template Layout",
	})
}

func TestTemplateLayout(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
