package application

import "context"

type PaymentIntentRecord struct {
	ID string
}

type PaymentIntentStore interface {
	Save(
		ctx context.Context,
		record PaymentIntentRecord,
	) error

	Get(
		ctx context.Context,
		id string,
	) (PaymentIntentRecord, bool, error)
}
