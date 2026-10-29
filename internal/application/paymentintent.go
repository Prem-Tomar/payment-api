package application

import (
	"context"
	"errors"
)

var ErrNotImplemented = errors.New("Payment creation intent is not implemented")

type CreatePaymentIntentInput struct {
	MerchantRef string
	CustomerRef string
	Amount      int
}

type CreatePaymentIntentResult struct {
	ID      string
	Status  PaymentIntentState
	Message string
}

type CreatePaymentIntentUseCase interface {
	CreatePaymentIntent(ctx context.Context, input CreatePaymentIntentInput) (CreatePaymentIntentResult, error)
}

type PaymentIntentService struct {
}

func (PaymentIntentService) CreatePaymentIntent(ctx context.Context, input CreatePaymentIntentInput) (CreatePaymentIntentResult, error) {

	if err := ctx.Err(); err != nil {
		return CreatePaymentIntentResult{}, err
	}
	_ = input

	return CreatePaymentIntentResult{
		ID:      "pi_placeholder_001",
		Status:  PaymentIntentStateCreated,
		Message: "Payment intent creation is not implemented.",
	}, nil
}

type FakePaymentIntentService struct{}

func (FakePaymentIntentService) CreatePaymentIntent(ctx context.Context, input CreatePaymentIntentInput) (CreatePaymentIntentResult, error) {
	if err := ctx.Err(); err != nil {
		return CreatePaymentIntentResult{}, err
	}

	return CreatePaymentIntentResult{
		ID:      "fake-payment-intent-001",
		Status:  PaymentIntentStateCreated,
		Message: "Fake payment intent result.",
	}, nil
}
