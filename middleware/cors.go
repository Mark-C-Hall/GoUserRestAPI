package middleware

import (
	"net/http"
	"strings"
	"user-api/config"
)

// CORSMiddleware is a middleware function that checks and sets CORS (Cross-Origin Resource Sharing)
// headers for incoming HTTP requests. It uses the allowed origins from the application's
// configuration to determine whether to set the CORS headers for the request's origin.
// If the request's method is OPTIONS (pre-flight CORS check), it will respond with a 200 OK status.
// Otherwise, it will call the next handler in the chain.
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		// Check if the request's origin is allowed and set the CORS headers accordingly
		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		}

		// If the request is an OPTIONS method (pre-flight CORS request), respond with 200 OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	}
}

// isOriginAllowed checks if the provided origin is in the list of allowed origins from the
// application's configuration. The function returns true if the origin is allowed, otherwise false.
// If the configuration allows all origins (using "*"), it will return true for any origin.
func isOriginAllowed(o string) bool {
	allowedOrigins := strings.Split(config.C.AllowedOrigins, ",")
	if allowedOrigins[0] == "*" {
		return true
	}
	for _, allowed := range allowedOrigins {
		if strings.EqualFold(o, allowed) {
			return true
		}
	}
	return false
}
