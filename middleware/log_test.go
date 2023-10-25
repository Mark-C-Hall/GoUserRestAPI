package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggingMiddleware(t *testing.T) {
	// Mock handler to simulate HTTP request processing
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			t.Fatalf("Could not write response: %v", err)
		}
	})

	// Wrap the mockHandler with the LoggingMiddleware
	handlerWithMiddleware := LoggingMiddleware(mockHandler)

	req, err := http.NewRequest("GET", "/test-endpoint", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Record the HTTP response using httptest.ResponseRecorder
	rr := httptest.NewRecorder()
	handlerWithMiddleware.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
