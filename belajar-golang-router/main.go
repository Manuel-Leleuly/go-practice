package main

import (
	"belajar-golang-router/helpers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    helpers.GetBaseUrl(false),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
