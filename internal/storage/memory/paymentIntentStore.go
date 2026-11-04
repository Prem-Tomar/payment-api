package memory

import (
	"context"
	"sync"

	"github.com/Prem-Tomar/payment-api/internal/application"
)

type paymentIntentKey struct {
	merchantID string
	id         string
}

type PaymentIntentStore struct {
	mu      sync.RWMutex
	records map[paymentIntentKey]application.PaymentIntentRecord
}

func NewPaymentIntentStore() *PaymentIntentStore {
	return &PaymentIntentStore{
		records: make(
			map[paymentIntentKey]application.PaymentIntentRecord,
		),
	}
}

func (s *PaymentIntentStore) Save(
	ctx context.Context,
	record application.PaymentIntentRecord,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	key := paymentIntentKey{
		merchantID: record.MerchantID,
		id:         record.ID,
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.records[key] = record
	return nil
}

func (s *PaymentIntentStore) Get(
	ctx context.Context,
	merchantID string,
	id string,
) (application.PaymentIntentRecord, bool, error) {
	if err := ctx.Err(); err != nil {
		return application.PaymentIntentRecord{}, false, err
	}
	key := paymentIntentKey{
		merchantID: merchantID,
		id:         id,
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	record, ok := s.records[key]
	return record, ok, nil
}
