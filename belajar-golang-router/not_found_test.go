package main

import (
	"belajar-golang-router/helpers"
	"fmt"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Gak Ketemu")
	})

	bytes := helpers.RunRouterTestDefault(router, "/")
	assert.Equal(t, "Gak Ketemu", string(bytes))
}
