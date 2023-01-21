package http

import "net/http"

func EnsureHandlerMethod(method string, h http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case method:
			h.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}
