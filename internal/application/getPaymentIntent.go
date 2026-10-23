package application

import (
	"context"
	"errors"
)

var ErrPaymentIntentNotFound = errors.New("payment intent not found")

type GetPaymentIntentResult struct {
	ID string
}

type GetPaymentIntentUseCase interface {
	GetPaymentIntent(ctx context.Context, id string) (GetPaymentIntentResult, error)
}

type PaymentIntentQueryService struct {
	store PaymentIntentStore
}

func NewPaymentIntentQueryService(store PaymentIntentStore) *PaymentIntentQueryService {
	return &PaymentIntentQueryService{
		store: store,
	}
}

func (s *PaymentIntentQueryService) GetPaymentIntent(ctx context.Context, id string) (GetPaymentIntentResult, error) {
	record, found, err := s.store.Get(ctx, id)
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
