package belajargolangweb

import (
	"belajar-golang-web/helpers"
	"embed"
	"html"
	"net/http"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	helpers.RunTemplateTest(TemplateCaching)
}

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]any{
		"Title": "Go-Lang Auto Escape",
		// "Body":  template.HTMLEscaper("<p>Selamat Belajar Go-Lang Web<script>alert('you've been hacked')</script></p>"),
		"Body": html.EscapeString("<p>Selamat Belajar Go-Lang Web<script>alert('you've been hacked')</script></p>"),
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	helpers.RunTemplateTest(TemplateAutoEscape)
}

func TestTemplateAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
