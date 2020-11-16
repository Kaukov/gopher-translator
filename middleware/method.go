package middleware

import (
	"net/http"
)

// NewMethod - adds a middleware to allow only a single HTTP method
// to be used for the endpoint
func NewMethod(method string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				rw.WriteHeader(http.StatusMethodNotAllowed)

				return
			}

			next.ServeHTTP(rw, r)
		})
	}
}
