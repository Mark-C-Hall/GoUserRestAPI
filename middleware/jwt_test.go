package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-api/util"
)

func TestJWTMiddleware(t *testing.T) {
	// Mock handler to simulate HTTP request processing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			t.Fatalf("Error writing response: %v", err)
		}
	})

	// Wrap the mockHandler with the JWTMiddleware
	handlerWithMiddleware := JWTMiddleware(mockHandler)

	tests := []struct {
		headerValue string
		blacklisted bool
		validToken  bool
		statusCode  int
	}{
		{"", false, false, http.StatusUnauthorized},
		{"Bearer ", false, false, http.StatusUnauthorized},
		{"Bearer blacklistedToken", true, true, http.StatusUnauthorized},
		{"Bearer invalidToken", false, false, http.StatusUnauthorized},
		{"Bearer validToken", false, true, http.StatusOK},
	}

	// Mock functions for the purpose of testing
	isTokenBlacklisted = func(token string) bool {
		return token == "blacklistedToken"
	}

	validateToken = func(token string) (*util.Claims, error) {
		if token == "validToken" {
			return &util.Claims{Username: "username"}, nil
		}
		return nil, errors.New("invalid token")
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/test-endpoint", nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		req.Header.Set("Authorization", test.headerValue)

		rr := httptest.NewRecorder()
		handlerWithMiddleware.ServeHTTP(rr, req)

		if rr.Code != test.statusCode {
			t.Errorf("Expected status code %v, but got %v for header value %v", test.statusCode, rr.Code, test.headerValue)
		}
	}
}

func TestExtractTokenFromRequest(t *testing.T) {
	tests := []struct {
		headerValue string
		expectToken string
		expectErr   bool
	}{
		{"", "", true},
		{"Bearer ", "", true},
		{"Bearer validToken", "validToken", false},
		{"InvalidPrefix validToken", "", true},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/test-endpoint", nil)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		req.Header.Set("Authorization", test.headerValue)

		token, err := extractTokenFromRequest(req)
		if test.expectErr && err == nil {
			t.Errorf("Expected error for header value %v, but got none", test.headerValue)
		} else if !test.expectErr && err != nil {
			t.Errorf("Didn't expect error for header value %v, but got: %v", test.headerValue, err.Error())
		} else if token != test.expectToken {
			t.Errorf("Expected token %v, but got %v for header value %v", test.expectToken, token, test.headerValue)
		}
	}
}
