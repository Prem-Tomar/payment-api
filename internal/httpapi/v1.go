package httpapi

import (
	"github.com/Prem-Tomar/payment-api/internal/application"
	"github.com/Prem-Tomar/payment-api/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func registerV1Routes(router *gin.Engine,
	createUseCase application.CreatePaymentIntentUseCase,
	getUseCase application.GetPaymentIntentUseCase) {
	v1 := router.Group("/v1")

	v1.Use(requestTimeOut(publicRequestTimeOut))
	v1.Use(requestBodyLimit())
	v1.Use(middlewares.RequireAPIKey)
	v1.Use(middlewares.RequireMerchantScope)

	rateLimiter := middlewares.NewRateLimiter()
	v1.Use(rateLimiter.Middleware())

	v1.POST("/payment-intents", requireIdempotencyKey, createPaymentIntentHandler(createUseCase))
	v1.GET("/payment-intents/:id", getPaymentIntentHandler(getUseCase))
}
