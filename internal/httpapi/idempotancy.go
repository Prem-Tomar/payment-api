package httpapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
  idempotencyKeyHeader = "Idempotency-Key"
	maxIdempotancyKeyLen = 255
)

func requireIdempotancyKey(context *gin.Context) {
	key := strings.TrimSpace(context.GetHeader(idempotencyKeyHeader))

	if key == "" {
		writeError(context, http.StatusBadRequest, "missing idempotency key")
		context.Abort()
		return
	}

	if len(key) > maxIdempotancyKeyLen {
		writeError(context, http.StatusBadRequest, "invalid idempotency key")
		context.Abort()
		return
	}

	context.Set("idempotency_key", key)
	context.Next()
}
