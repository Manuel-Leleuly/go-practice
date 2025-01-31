package main

import (
	"belajar-golang-router/helpers"
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id") + " Item " + p.ByName("itemId")
		fmt.Fprint(w, text)
	})

	url := helpers.CreateUrl("/products/1/items/2", nil)
	request := helpers.CreateRequest(http.MethodGet, url, nil)
	bytes := helpers.RunRouterTest(router, request)
	assert.Equal(t, "Product 1 Item 2", string(bytes))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Image: " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	bytes := helpers.RunRouterTestDefault(router, "/images/small/profile.png")
	assert.Equal(t, "Image: /small/profile.png", string(bytes))
}
