package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"text/template"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Manuel",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]any{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateComparator(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]any{
		"Title":   "Template Action Range",
		"Hobbies": []string{"Gaming", "Reading", "Coding"},
	})
}

func TestTemplateRange(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateAddress(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]any{
		"Title": "Template Action Address",
		"Name":  "Manuel",
		"Address": map[string]any{
			"Street": "Jalan Belum Ada",
			"City":   "Jakarta",
		},
	})
}

func TestTemplateAddress(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateAddress(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
