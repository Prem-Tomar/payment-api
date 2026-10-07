package httpapi

import "github.com/gin-gonic/gin"

type requestMetadata struct {
	RequestID      string
	IdempotencyKey string
}

func buildRequestMetadata(c *gin.Context) requestMetadata {
	requestID, _ := c.Get("request_id")

	requestIDValue, _ := requestID.(string)

	return requestMetadata{
		RequestID:      requestIDValue,
		IdempotencyKey: c.GetHeader("Idempotency-Key"),
	}
}
