package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"text/template"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]any{
		"Title": "Template Data Map",
		"Name":  "Manuel",
		"Address": map[string]any{
			"Street": "Jalan Belum Ada Lagi",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Manuel",
		Address: Address{
			Street: "Jalan belum ada",
		},
	})
}

func TestTemplateDatStruct(t *testing.T) {
	request, recorder := GetReqRec()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
