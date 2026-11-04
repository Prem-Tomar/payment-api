package application

import "context"

type PaymentIntentRecord struct {
	ID         string
	MerchantID string
}

type PaymentIntentStore interface {
	Save(
		ctx context.Context,
		record PaymentIntentRecord,
	) error

	Get(
		ctx context.Context,
		merchantID string,
		id string,
	) (PaymentIntentRecord, bool, error)
}
