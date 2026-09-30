package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	MerchantIDHeader = "X-Merchant-ID"
	MerchantIDKey    = "merchant_id"
)

func RequireMerchantScope(c *gin.Context) {
	merchantID := strings.TrimSpace(c.GetHeader(MerchantIDHeader))

	if merchantID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing merchant scope"})
		return
	}

	if strings.ContainsAny(merchantID, " \t\r\n") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid merchant scope"})
		return
	}
	c.Set(MerchantIDKey, merchantID)
	c.Next()
}
