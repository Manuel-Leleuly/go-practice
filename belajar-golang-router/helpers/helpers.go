package helpers

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"
)

func GetBaseUrl(withProtocol bool) string {
	const address string = "localhost:3000"
	if withProtocol {
		return "http://" + address
	}
	return address
}

func CreateUrl(path string, queryParams map[string]string) string {
	params := ""
	for k, v := range queryParams {
		if params == "" {
			params = k + "=" + v
		} else {
			params += "&" + k + "=" + v
		}
	}
	return GetBaseUrl(true) + path + "?" + params
}

func CreateRequest(method string, url string, requestBody io.Reader) *http.Request {
	return httptest.NewRequest(method, url, requestBody)
}

func RunRouterTest(router *httprouter.Router, request *http.Request) []byte {
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	return bytes
}

// Don't use this if everything is custom
func RunRouterTestDefault(router *httprouter.Router, urlPath string) []byte {
	url := CreateUrl(urlPath, nil)
	request := CreateRequest(http.MethodGet, url, nil)
	return RunRouterTest(router, request)
}
