package application

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("Payment creation intent is not implemented")

type CreatePaymentIntentResult struct {
	ID string
}

type CreatePaymentIntentUseCase interface {
	CreatePaymentIntent(ctx context.Context) (CreatePaymentIntentResult, error)
}

type PaymentIntentService struct {
}

func (PaymentIntentService) CreatePaymentIntent(ctx context.Context) (CreatePaymentIntentResult, error) {

	return CreatePaymentIntentResult{}, ErrNotImplemented
}
