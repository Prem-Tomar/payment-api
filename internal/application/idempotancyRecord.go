package application

import (
	"time"
)

type IdempotencyRecord struct {
	Key         string
	Operation   string
	RequestHash string
	ResourceID  string
	CreatedAt   time.Time
}

type IdempotencyRecordStore interface {
	Get(
		key string,
		operation string,
	) (IdempotencyRecord, bool, error)

	Save(
		record IdempotencyRecord,
	) error
}
