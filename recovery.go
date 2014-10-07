// Package recovery implements a simple net/http recovery middleware
package recovery

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
)

type recoveryHandler struct {
	writer     io.Writer
	handler    http.Handler
	PrintStack bool
}

// Handler returns a new http.Handler that will wrap another handler and recover
// any panics
func Handler(out io.Writer, h http.Handler, printStack bool) http.Handler {
	return recoveryHandler{writer: out, handler: h}
}

func (h recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			stack := debug.Stack()
			f := "PANIC: %s\n%s"
			fmt.Fprintf(h.writer, f, err, stack)
			if h.PrintStack {
				fmt.Fprintf(w, f, err, stack)
			}
		}
	}()
	h.handler.ServeHTTP(w, r)
}
