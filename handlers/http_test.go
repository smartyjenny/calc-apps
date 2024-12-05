package handlers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func assertHTTP(t *testing.T, method, target string, expectedStatus int, expectedContentType, expectedResponse string) {
	request := httptest.NewRequest(method, target, nil)
	response := httptest.NewRecorder()

	NewRouter(nil).ServeHTTP(response, request)

	assertEqual(t, response.Body.String(), expectedResponse)
	assertEqual(t, response.Code, expectedStatus)
	assertEqual(t, response.Header().Get("Content-Type"), expectedContentType)
}

func TestHTTPServer_Add(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/add?a=1&b=2", http.StatusOK, "text/plain; charset=utf-8", "3")
	assertHTTP(t, http.MethodPost, "/add?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/add?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid")
}

func assertEqual(t *testing.T, actual, expected any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("want [%v], got [%v]\n", expected, actual)
	}
}
