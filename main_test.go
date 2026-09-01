package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// the below are table tests
	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "returns healthy response",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/healthz", nil)   // create a new fake request
			rec := httptest.NewRecorder()                                 // records this response

			healthHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.expectedStatus,
					rec.Code,
				)
			}

			if rec.Body.String() != tt.expectedBody+"\n" {
				t.Fatalf(
					"expected body %q, got %q",
					tt.expectedBody+"\n",
					rec.Body.String(),
				)
			}
		})
	}
}

func TestReadyHandler(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "returns ready response",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ready"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
			rec := httptest.NewRecorder()

			readyHandler(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.expectedStatus,
					rec.Code,
				)
			}

			if rec.Body.String() != tt.expectedBody+"\n" {
				t.Fatalf(
					"expected body %q, got %q",
					tt.expectedBody+"\n",
					rec.Body.String(),
				)
			}
		})
	}
}
