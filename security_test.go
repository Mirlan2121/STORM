package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler := SecurityMiddleware(
		LoggingMiddleware(
			http.HandlerFunc(handleRequest),
		),
	)

	handler.ServeHTTP(rr, req)

	// Проверяем security headers
	headers := map[string]string{
		"X-XSS-Protection":       "1; mode=block",
		"X-Frame-Options":        "DENY",
		"X-Content-Type-Options": "nosniff",
	}

	for header, expected := range headers {
		if got := rr.Header().Get(header); got != expected {
			t.Errorf("%s header = %v, want %v", header, got, expected)
		}
	}
}
