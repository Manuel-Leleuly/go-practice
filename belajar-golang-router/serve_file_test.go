package main

import (
	"belajar-golang-router/helpers"
	"embed"
	"io/fs"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	bytes := helpers.RunRouterTestDefault(router, "/files/hello.txt")
	assert.Equal(t, "Hello HttpRouter", string(bytes))
}
