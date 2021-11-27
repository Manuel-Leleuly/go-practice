package belajargolangweb

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	/*
		prev code: mux.Handle("/static/", fileServer)

		When we run the test using the prev code, and type "localhost:8080/index.html" or .css or .js, the result is always be "error 404 not found".
		This is because the handler set by mux will find the file inside the folder "static".
		For example, if we type "localhost/index.js", the server will try to search in "/resources/static/".

		The solution to that is to use StripPrefix() method to remove the prefix of the file source

	*/

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
	uploading both resources and static file manually can be quite a hassle.
	Therefore, we can use embed to embed file into a binary distribution file
*/

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")
	fileServer := http.FileServer(http.FS(directory))

	/*
		prev code: fileServer := http.FileServer(http.FS(resources))

		when you run the test and type "localhost:8080/static/", the result is the "resources" folder.
		This is because embed will consider the folder name as the path.
		In order to avoid this, we can use fs.Sub()
	*/

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
