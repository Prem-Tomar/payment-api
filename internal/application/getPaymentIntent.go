package application

import (
	"context"
	"errors"
)

var ErrPaymentIntentNotFound = errors.New("payment intent not found")

type GetPaymentIntentResult struct {
	ID string
}
type GetPaymentIntentInput struct {
	ID         string
	MerchantID string
}

type GetPaymentIntentUseCase interface {
	GetPaymentIntent(ctx context.Context, input GetPaymentIntentInput) (GetPaymentIntentResult, error)
}

type PaymentIntentQueryService struct {
	store PaymentIntentStore
}

func NewPaymentIntentQueryService(store PaymentIntentStore) *PaymentIntentQueryService {
	return &PaymentIntentQueryService{
		store: store,
	}
}

func (s *PaymentIntentQueryService) GetPaymentIntent(ctx context.Context, input GetPaymentIntentInput) (GetPaymentIntentResult, error) {
	record, found, err := s.store.Get(ctx, input.MerchantID, input.ID)
	if err != nil {
		return GetPaymentIntentResult{}, err
	}
	if !found {
		return GetPaymentIntentResult{}, ErrPaymentIntentNotFound
	}
	return GetPaymentIntentResult{
		ID: record.ID,
	}, nil

}
