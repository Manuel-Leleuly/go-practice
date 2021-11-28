package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak Boleh")
	})

	router.POST("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "POST")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Boleh", string(body))
}
