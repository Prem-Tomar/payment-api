package httpapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Prem-Tomar/payment-api/internal/application"
	"github.com/Prem-Tomar/payment-api/internal/httpapi/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
					ErrRequestTooLarge,
					"Request body to large",
				)
				return
			}
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				writeError(
					c,
					http.StatusBadRequest,
					ErrValidation,
					"Request field failed valiation.",
				)
				return
			}

			writeError(
				c,
				http.StatusBadRequest,
				ErrInvalidRequest,
				"invalid request body",
			)
			return

		}

		request.MerchantRef = strings.TrimSpace(request.MerchantRef)
		request.CustomerRef = strings.TrimSpace(request.CustomerRef)

		if request.MerchantRef == "" {
			writeError(
				c,
				http.StatusBadRequest,
				ErrInvalidRequest,
				"Marchant reference is required",
			)
			return
		}

		metadata := buildRequestMetadata(c)
		_ = metadata

		input := application.CreatePaymentIntentInput{
			MerchantRef: request.MerchantRef,
			CustomerRef: request.CustomerRef,
			Amount:      int(request.Amount),
		}

		result, err := useCase.CreatePaymentIntent(c.Request.Context(), input)
		if errors.Is(err, application.ErrNotImplemented) {
			writeError(
				c,
				http.StatusNotImplemented,
				ErrNotImplemented,
				"payment intent creation is not implemented",
			)
			return
		}

		if err != nil {
			writeError(
				c,
				http.StatusInternalServerError,
				ErrInternal,
				"internal server error",
			)
			return
		}

		response := dto.CreatePaymentIntentResponse{
			Status: dto.PaymentIntentStatus(result.Status),
			ID:     result.ID,
		}

		c.JSON(http.StatusCreated, response)

	}
}

func getPaymentIntentHandler(c *gin.Context) {
	id := strings.TrimSpace(c.Param("id"))

	if id == "" {
		writeError(c, http.StatusBadRequest, ErrValidation, "payment intent id is required")
		return
	}

	writeError(
		c, http.StatusNotImplemented, ErrNotImplemented, "payment intent lookup is not implemented",
	)
}
