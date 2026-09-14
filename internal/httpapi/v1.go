package httpapi

import (
	"net/http"

	"github.com/Prem-Tomar/payment-api/internal/httpapi/dto"
	"github.com/gin-gonic/gin"
)

func registerV1Routes(router *gin.Engine) {
	v1 := router.Group("/v1")

	v1.POST("/payment-intents", requireIdempotancyKey, createPaymentIntentHandler)
}

func createPaymentIntentHandler(context *gin.Context) {
	var request dto.CreatePaymentIntentRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		writeError(context, http.StatusBadRequest, "invalid request body")
		return
	}
}
