package httpapi

import (
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
	if want := `{"error":"method not allowed"}` + "\n"; rec.Body.String() != want {
		t.Fatalf("expected body %q, got %q", want, rec.Body.String())
	}
}

func TestBodyLimitRejectsDeclaredOversizedRequest(t *testing.T) {
	router := testRouter()
	body := strings.NewReader(strings.Repeat("a", maxRequestBodySize+1))
	req := httptest.NewRequest(http.MethodPost, "/healthz", body)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("expected status %d, got %d", http.StatusRequestEntityTooLarge, rec.Code)
	}
	if want := `{"error":"request body too large"}` + "\n"; rec.Body.String() != want {
		t.Fatalf("expected body %q, got %q", want, rec.Body.String())
	}
}
