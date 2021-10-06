package storage

import "context"

type VisitCounterRepository interface {
	AggregateCounters(ctx context.Context) error
	Close()
}
