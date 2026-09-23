package application

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("Payment creation intent is not implemented")

type CreatePaymentIntentUseCase interface {
	CreatePaymentIntent(ctx context.Context) error
}

type PaymentIntentService struct {
	Status string `json:"status"`
}

func (PaymentIntentService) CreatePaymentIntent(ctx context.Context) error {
	return ErrNotImplemented
}
