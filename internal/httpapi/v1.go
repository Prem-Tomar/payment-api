package httpapi

import (
	"errors"
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
