package storage

import "context"

type VisitCounterRepository interface {
	Increase(ctx context.Context, counter string) error
	Close()
}
