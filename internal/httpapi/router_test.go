package httpapi

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewRouter(logger)
}

func TestOperationalEndpoints(t *testing.T) {
	tests := []struct {
		name string
		path string
		body string
	}{
		{name: "health", path: "/healthz", body: `{"status":"ok"}`},
		{name: "readiness", path: "/readyz", body: `{"status":"ready"}`},
	}

	router := testRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
			}
			if rec.Body.String() != tt.body {
				t.Fatalf("expected body %q, got %q", tt.body, rec.Body.String())
			}
		})
	}
}

func TestRequestIDMiddleware(t *testing.T) {
	router := testRouter()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	req.Header.Set("X-Request-ID", "req-123")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if got := rec.Header().Get("X-Request-ID"); got != "req-123" {
		t.Fatalf("expected echoed request id req-123, got %q", got)
	}
}

func TestMethodNotAllowedResponse(t *testing.T) {
	router := testRouter()
	req := httptest.NewRequest(http.MethodPost, "/healthz", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
	if got := rec.Header().Get("Allow"); got != "GET" {
		t.Fatalf("expected Allow header GET, got %q", got)
	}
	if want := `{"code":"METHOD_NOT_ALLOWED","error":"method not allowed"}`; rec.Body.String() != want {
		t.Fatalf("expected body %q, got %q", want, rec.Body.String())
	}
}

func TestBodyLimitRejectsDeclaredOversizedRequest(t *testing.T) {
	router := testRouter()
	body := strings.NewReader(strings.Repeat("a", maxRequestBodySize+1))
	req := httptest.NewRequest(http.MethodPost, "/v1/payment-intents", body)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("expected status %d, got %d", http.StatusRequestEntityTooLarge, rec.Code)
	}
	if want := `{"code":"REQUEST_TOO_LARGE","error":"request body too large"}`; rec.Body.String() != want {
		t.Fatalf("expected body %q, got %q", want, rec.Body.String())
	}
}

func TestPaymentIntentErrorCodes(t *testing.T) {
	tests := []struct {
		name string
		body string
		code string
	}{
		{name: "malformed JSON", body: `{"amount":`, code: "INVALID_REQUEST"},
		{name: "invalid amount", body: `{"amount":0,"currency":"USD","payment_method":"card","merchant_ref":"order-1"}`, code: "VALIDATION_ERR"},
		{name: "blank merchant reference", body: `{"amount":100,"currency":"USD","payment_method":"card","merchant_ref":"   "}`, code: "VALIDATION_ERR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := testRouter()
			req := httptest.NewRequest(http.MethodPost, "/v1/payment-intents", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-API-Key", "test-key")
			req.Header.Set("X-Merchant-ID", "merchant-1")
			req.Header.Set("Idempotency-Key", "request-1")
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)
			if rec.Code != http.StatusBadRequest {
				t.Fatalf("expected status 400, got %d: %s", rec.Code, rec.Body.String())
			}
			var response struct {
				Code string `json:"code"`
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if response.Code != tt.code {
				t.Fatalf("expected code %q, got %q", tt.code, response.Code)
			}
		})
	}
}
