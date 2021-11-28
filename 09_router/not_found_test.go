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

func TestNotFound(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Gak Ketemu")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/404", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(body))
}
