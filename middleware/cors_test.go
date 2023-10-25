package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"user-api/config"
)

func TestCORSMiddleware(t *testing.T) {
	// Mock handler to simulate HTTP request processing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			t.Fatalf("Error writing response: %v", err)
		}
	})

	// Wrap the mockHandler with the CORSMiddleware
	handlerWithMiddleware := CORSMiddleware(mockHandler)

	// Define test cases
	tests := []struct {
		origin      string
		expectAllow bool
	}{
		{"http://allowed-origin.com", true},
		{"http://disallowed-origin.com", false},
		{"", false},
	}

	// Mock allowed origins for testing
	config.C.AllowedOrigins = "http://allowed-origin.com"

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/test-endpoint", nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		req.Header.Set("Origin", test.origin)

		rr := httptest.NewRecorder()
		handlerWithMiddleware.ServeHTTP(rr, req)

		if test.expectAllow {
			if rr.Header().Get("Access-Control-Allow-Origin") != test.origin {
				t.Errorf("Expected allowed origin %v but got %v", test.origin, rr.Header().Get("Access-Control-Allow-Origin"))
			}
		} else {
			if rr.Header().Get("Access-Control-Allow-Origin") != "" {
				t.Errorf("Expected no allowed origin but got %v", rr.Header().Get("Access-Control-Allow-Origin"))
			}
		}
	}
}

func TestIsOriginAllowed(t *testing.T) {
	// Mock allowed origins for testing
	config.C.AllowedOrigins = "http://allowed-origin.com,http://another-allowed.com"

	tests := []struct {
		origin  string
		allowed bool
	}{
		{"http://allowed-origin.com", true},
		{"http://disallowed-origin.com", false},
		{"http://another-allowed.com", true},
		{"", false},
	}

	for _, test := range tests {
		result := isOriginAllowed(test.origin)
		if result != test.allowed {
			t.Errorf("Expected origin %v to be allowed: %v but got %v", test.origin, test.allowed, result)
		}
	}
}
