package httpapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Prem-Tomar/payment-api/internal/application"
	"github.com/Prem-Tomar/payment-api/internal/httpapi/dto"
	"github.com/gin-gonic/gin"
)

func registerV1Routes(router *gin.Engine, useCase application.CreatePaymentIntentUseCase) {
	v1 := router.Group("/v1")

	v1.POST("/payment-intents", requireIdempotencyKey, createPaymentintentHandler)
	 v1.GET(
        "/payment-intents/:id",
        getPaymentIntentHandler,
    )
}

func createPaymentintentHandler(context *gin.Context) {
	var request dto.CreatePaymentIntentRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		var maxBytesErr *http.MaxBytesError

		if errors.As(err, &maxBytesErr) {
			writeError(
				context,
				http.StatusRequestEntityTooLarge,
				"request body too large",
			)
			return
		}

		writeError(
			context,
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}

	writeError(
		context,
		http.StatusNotImplemented,
		"payment intent creation is not implemented",
	)
}

func getPaymentIntentHandler(c *gin.Context){
	id := strings.TrimSpace(c.Param("id"))

	if id =="" {
		writeError(c, http.StatusBadRequest, "payment intent id is required")
		return 
	}

	_ = id 
	writeError(
		c, http.StatusNotImplemented, "payment intent lookup is not implemented",
	)
}