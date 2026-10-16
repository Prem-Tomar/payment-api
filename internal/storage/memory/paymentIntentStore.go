package memory

import (
	"context"
	"sync"

	"github.com/Prem-Tomar/payment-api/internal/application"
)

type PaymentIntentStore struct {
	mu      sync.RWMutex
	records map[string]application.PaymentIntentRecord
}

func NewPaymentIntentStore() *PaymentIntentStore {
	return &PaymentIntentStore{
		records: make(
			map[string]application.PaymentIntentRecord,
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

	s.mu.Lock()
	defer s.mu.Unlock()

	s.records[record.ID] = record
	return nil
}

func (s *PaymentIntentStore) Get(
	ctx context.Context,
	id string,
) (application.PaymentIntentRecord, bool, error) {
	if err := ctx.Err(); err != nil {
		return application.PaymentIntentRecord{}, false, err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	record, ok := s.records[id]
	return record, ok, nil
}
