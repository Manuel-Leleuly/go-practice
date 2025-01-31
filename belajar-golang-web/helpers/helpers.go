package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

var BaseUrl = "http://localhost:8080"

func GetReqRec() (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(http.MethodGet, BaseUrl, nil)
	recorder := httptest.NewRecorder()
	return request, recorder
}

func RunTemplateTest(handlerFunction func(http.ResponseWriter, *http.Request)) {
	request, recorder := GetReqRec()
	handlerFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
