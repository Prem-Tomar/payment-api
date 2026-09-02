package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// the below are table tests
	tests := []struct {
		// This is the structure of test case
		name           string
		method         string
		expectedStatus int
		expectedBody   string
	}{
		{ // trst case 1
			name:           "returns healthy response",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"ok"}`,
		},
		{ // Test case 2
			name:           "rejects non GET request",
			method:         http.MethodPost,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	router := newRouter()

	for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		req := httptest.NewRequest(tt.method, "/healthz", nil)
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		if rec.Code != tt.expectedStatus {
			t.Fatalf(
				"expected status %d, got %d",
				tt.expectedStatus,
				rec.Code,
			)
		}

		if tt.expectedBody != "" &&
			rec.Body.String() != tt.expectedBody+"\n" {
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
