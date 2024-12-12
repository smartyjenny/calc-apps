package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	calc_lib "github.com/smartyjenny/calc-lib"
)

type Router struct {
	logger log.Logger
}

func NewRouter(logger *log.Logger) http.Handler {
	router := http.NewServeMux()
	router.Handle("GET /add", newHTTPHandler(logger, &calc_lib.Addition{}))
	router.Handle("GET /sub", newHTTPHandler(logger, &calc_lib.Subtraction{}))
	router.Handle("GET /multiply", newHTTPHandler(logger, &calc_lib.Multiplication{}))
	router.Handle("GET /division", newHTTPHandler(logger, &calc_lib.Division{}))
	return router
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator Calculator
}

func newHTTPHandler(logger *log.Logger, calculator Calculator) http.Handler {
	return &HTTPHandler{logger: logger, calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		response.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprint(response, "a was invalid")
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		response.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = fmt.Fprint(response, "b was invalid")
		return
	}

	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(response, c)
	if err != nil {
		log.Fatal(err)
	}

}
