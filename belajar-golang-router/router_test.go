package main

import (
	"belajar-golang-router/helpers"
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	url := helpers.CreateUrl("/", nil)
	request := helpers.CreateRequest(http.MethodGet, url, nil)
	bytes := helpers.RunRouterTest(router, request)
	assert.Equal(t, "Hello World", string(bytes))
}

func TestParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id")
		fmt.Fprint(w, text)
	})

	bytes := helpers.RunRouterTestDefault(router, "/products/1")
	assert.Equal(t, "Product 1", string(bytes))
}
