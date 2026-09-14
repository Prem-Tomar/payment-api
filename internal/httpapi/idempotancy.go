package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	idempotencyKeyHeader = "Idempotancy-key"
	maxIdempotancyKeyLen = 255
)

func requireIdempotancyKey(context *gin.Context) {
	key := context.GetHeader(idempotencyKeyHeader)

	if key == "" {
		writeError(context, http.StatusBadRequest, "missing idempotancy key")
		context.Abort()
		return
	}

	if len(key) > maxIdempotancyKeyLen {
		writeError(context, http.StatusBadRequest, "invalid idempotancy key")
		context.Abort()
		return
	}

	context.Set("idempotancy_key", key)
	context.Next()
}
