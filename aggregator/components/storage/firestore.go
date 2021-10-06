package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type FirestoreVisitCounterRepositoryConfig struct {
	ProjectId string
}

type firestoreVisitCounterRepository struct {
	config *FirestoreVisitCounterRepositoryConfig
	client *firestore.Client
}

func CreateFirestoreVisitCounterRepository(
	ctx context.Context,
	config *FirestoreVisitCounterRepositoryConfig) (VisitCounterRepository, error) {

	var err error
	var client *firestore.Client

	// Initialize Firestore client.
	if client, err = firestore.NewClient(ctx, config.ProjectId); err != nil {
		return nil, err
	}

	return &firestoreVisitCounterRepository{config, client}, nil
}

func (repo *firestoreVisitCounterRepository) Close() {
	repo.client.Close()
}

func (repo *firestoreVisitCounterRepository) AggregateCounters(ctx context.Context) error {
	// Counters are stored in the Firestore database according to the following structure:
	// /counters/visits/page/shard

	// Get a reference to the counters collection.
	counters := repo.client.Collection("counters")
	// Get a reference to the visits document.
	visits := counters.Doc("visits")
	// Get all collections under visits. There is one for each page.
	pages := visits.Collections(ctx)

	var total int64 = 0

	for {
		page, err := pages.Next()

		// Handle the error.
		if err != nil && err != iterator.Done {
			return err
		} else if err == iterator.Done {
			break
		}

		// The current collection holds a set of shards for the counter in question.
		shards := page.Documents(ctx)

		for {
			// Get the next shard.
			shard, err := shards.Next()

			if err != nil && err != iterator.Done {
				return err
			} else if err == iterator.Done {
				break
			}
			// Get the count from the shard.
			data, found := shard.Data()["Count"]

			if found {
				total += data.(int64)
			} else {
				return fmt.Errorf("unexpected data format in %s", shard.Ref.Path)
			}
		}
	}

	// We have the total count now; we will write the aggregate result to reduce the
	// amount of requests that need to be made when reading.
	// The aggregate count is written to /counters/aggregates
	fmt.Println("Setting total visits to be", total)

	_, err := counters.Doc("aggregates").Set(ctx, map[string]interface{}{
		"visits": total,
	})
	return err
}
