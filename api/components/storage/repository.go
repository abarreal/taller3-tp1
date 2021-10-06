package storage

import "context"

type VisitCounterRepository interface {
	GetCount(ctx context.Context) (uint64, error)
	Close()
}
