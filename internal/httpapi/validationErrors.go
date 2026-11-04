package httpapi

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func mapValidationError(err error) (APIErrorCode, string, bool) {
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return "", "", false
	}

	if len(validationErrors) == 0 {
		return "", "", false
	}

	field := strings.ToLower(validationErrors[0].Field())

	switch field {
	case "amount":
		return ErrInvalidAmount, "amount must be greater than zero", true

	case "currency":
		return ErrInvalidCurrency, "currency is required and must contain exactly 3 English letters", true

	case "paymentmethod":
		return ErrInvalidPaymentMethod, "payment_method is required or invalid", true

	case "merchantref", "customerref":
		return ErrInvalidReference, "merchant_ref or customer_ref is invalid", true

	default:
		return ErrValidation, "request validation failed", true
	}
}
