package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware function that logs the method, request URI,
// and the duration it took for the request to be processed.
// It takes as its parameter a function that conforms to the http.HandlerFunc type
// (i.e., a function that can be registered to serve a particular pattern in the Go HTTP package).
// The function returns another http.HandlerFunc. This allows for chaining or
// layering multiple middleware functions if needed.
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Record the start time of the request processing
		start := time.Now()

		// Call the next handler or middleware in the chain
		next.ServeHTTP(w, r)

		// Log the HTTP method, request URI, and the duration of the request
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}
