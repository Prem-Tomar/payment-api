package httpapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	idempotencyKeyHeader = "Idempotency-Key"
	maxIdempotencyKeyLen = 255
)

func requireIdempotencyKey(context *gin.Context) {
	key := strings.TrimSpace(context.GetHeader(idempotencyKeyHeader))

	if key == "" {
		writeError(context, http.StatusBadRequest,ErrMissingIdempotencyKey, "missing idempotency key")
		context.Abort()
		return
	}

	if len(key) > maxIdempotencyKeyLen {
		writeError(context, http.StatusBadRequest, ErrInvalidIdempotencyKey,"invalid idempotency key")
		context.Abort()
		return
	}

	context.Set("idempotency_key", key)
	context.Next()
}
