package httpapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Prem-Tomar/payment-api/internal/application"
	"github.com/Prem-Tomar/payment-api/internal/httpapi/dto"
	"github.com/gin-gonic/gin"
)

func createPaymentIntentHandler(useCase application.CreatePaymentIntentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.CreatePaymentIntentRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			var maxBytesErr *http.MaxBytesError

			if errors.As(err, &maxBytesErr) {
				writeError(
					c,
					http.StatusRequestEntityTooLarge,
					"Request body to large",
				)
				return
			}

			writeError(
				c,
				http.StatusBadRequest,
				"invalid request body",
			)
			return
		}

		result, err := useCase.CreatePaymentIntent(c.Request.Context())
		if errors.Is(err, application.ErrNotImplemented) {
			writeError(
				c,
				http.StatusNotImplemented,
				"payment intent creation is not implementd",
			)
			return
		}

		if err != nil {
			writeError(
				c,
				http.StatusInternalServerError,
				"internal server error",
			)
			return
		}

		response := dto.CreatePaymentIntentResponse{
			Status: "created",
			ID:     result.ID,
		}

		c.JSON(http.StatusCreated, response)

	}
}

func getPaymentIntentHandler(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))

	if id == "" {
		writeError(c, http.StatusBadRequest, "payment intent id is required")
		return
	}

	writeError(
		c, http.StatusNotImplemented, "payment intent lookup is not implemented",
	)
}
