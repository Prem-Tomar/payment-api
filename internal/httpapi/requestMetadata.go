package httpapi

import (
	"github.com/Prem-Tomar/payment-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type requestMetadata struct {
	RequestID      string
	IdempotencyKey string
	MerchantID     string
}

func buildRequestMetadata(c *gin.Context) requestMetadata {
	requestID, _ := c.Get("request_id")

	requestIDValue, _ := requestID.(string)

	merchantID, _ := c.Get(middlewares.MerchantIDKey)
	merchantIDValue, _ := merchantID.(string)

	return requestMetadata{
		RequestID:      requestIDValue,
		IdempotencyKey: c.GetHeader("Idempotency-Key"),
		MerchantID:     merchantIDValue,
	}
}
