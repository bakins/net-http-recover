package recovery

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newRequest(method, url string) *http.Request {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	return req
}

func TestHandler(t *testing.T) {
	var buf bytes.Buffer

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "Hello World\n")
	})
	r := Handler(&buf, handler, false)
	r.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/foo"))

	if buf.Len() > 0 {
		t.Fatal("buffer should be empty")
	}
}

func TestHandlerWithPanic(t *testing.T) {
	var buf bytes.Buffer

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		panic("something bad happened")
	})
	r := Handler(&buf, handler, false)
	r.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/foo"))

	if !strings.Contains(buf.String(), "something bad happened") {
		t.Fatal("buffer did not match expected result")
	}
}

func TestHandlerWithPanicPrint(t *testing.T) {
	var buf bytes.Buffer

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		panic("something bad happened")
	})
	r := Handler(&buf, handler, true)
	r.ServeHTTP(httptest.NewRecorder(), newRequest("GET", "/foo"))

	if !strings.Contains(buf.String(), "something bad happened") {
		t.Fatal("buffer did not match expected result")
	}
}
