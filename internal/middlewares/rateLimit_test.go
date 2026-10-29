package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRateLimitFirstRequestHeaders(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(func(c *gin.Context) {
		c.Set(MerchantIDKey, "merchant-a")
		c.Next()
	})
	router.Use(NewRateLimiter().Middleware())
	router.GET("/limited", func(c *gin.Context) { c.Status(http.StatusNoContent) })
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/limited", nil))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected first request to succeed, got %d", rec.Code)
	}
	if got := rec.Header().Get("X-RateLimit-Limit"); got != "5" {
		t.Fatalf("expected limit 5, got %q", got)
	}
	if got := rec.Header().Get("X-RateLimit-Remaining"); got != "4" {
		t.Fatalf("expected 4 remaining after one request, got %q", got)
	}
	if got := rec.Header().Get("Retry-After"); got != "" {
		t.Fatalf("expected no retry delay on allowed request, got %q", got)
	}
}
