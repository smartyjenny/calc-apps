package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend-training/cohort-c-2/calc-apps/externals/should"
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

func TestHTTPServer_Subtract(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/sub?a=2&b=1", http.StatusOK, "text/plain; charset=utf-8", "1")
	assertHTTP(t, http.MethodPost, "/sub?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/sub?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid")
	assertHTTP(t, http.MethodGet, "/sub?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid")
}

func TestHTTPServer_Multiply(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/multiply?a=2&b=1", http.StatusOK, "text/plain; charset=utf-8", "2")
	assertHTTP(t, http.MethodPost, "/multiply?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/multiply?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid")
	assertHTTP(t, http.MethodGet, "/multiply?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid")
}

func TestHTTPServer_Division(t *testing.T) {
	assertHTTP(t, http.MethodGet, "/division?a=4&b=2", http.StatusOK, "text/plain; charset=utf-8", "2")
	assertHTTP(t, http.MethodPost, "/division?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/division?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid")
	assertHTTP(t, http.MethodGet, "/division?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid")
}

func TestHTTPServer_NotFound(t *testing.T) {
	// bogus method
}

func assertEqual(t *testing.T, actual, expected any) {
	err := should.Equal(actual, expected)
	if err != nil {
		t.Errorf("expected [%v], got [%v]\n", expected, actual)
	}
}
