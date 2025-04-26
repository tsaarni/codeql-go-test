package codeqltest

import (
	"errors"
	"fmt"
	"net/http"
)

var lastRequestError *MyRequest

var errHelloWorld = errors.New("hello world")

type MyRequest struct {
	*http.Request
	headerString string
}

func (m *MyRequest) formatHeaders() string {
	var h string
	for k, v := range m.Request.Header {
		h = fmt.Sprintf("%s%s: %s\n", h, k, v)
	}
	return h
}

func (m *MyRequest) Error() string {
	m.headerString = fmt.Sprintf("Error: %s", m.formatHeaders())
	return "foo"
}

func SetLastRequestError(r *http.Request) {
	lastRequestError = &MyRequest{Request: r}
}

func GetLastRequestError() string {
	return lastRequestError.Error()
}

func RunWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		SetLastRequestError(r)
		fmt.Println(errHelloWorld.Error())
	})
	http.ListenAndServe(":8080", nil)
}
