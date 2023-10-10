package middleware

import (
	"net/http"
	"strings"
	"user-api/config"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if isOriginAllowed(origin) {
			// Set headers
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		}

		// If it's just an OPTIONS request (pre-flight) return with 200 OK
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler (the main one)
		next.ServeHTTP(w, r)
	}
}

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
