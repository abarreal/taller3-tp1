package storage

import (
	"context"
	"math/rand"
	"strconv"
	"strings"

	"cloud.google.com/go/firestore"
)

type FirestoreVisitCounterRepositoryConfig struct {
	ProjectId        string
	ShardsPerCounter uint32
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

func (repo *firestoreVisitCounterRepository) Increase(
	ctx context.Context, counter string) error {

	// Select a shard at random.
	idx := rand.Intn(int(repo.config.ShardsPerCounter))

	// The sharded counters are stored in /counters/visits as subcollections.
	colCounters := repo.client.Collection("counters")
	docVisits := colCounters.Doc("visits")
	// Construct a reference to the counter shard to be increased.
	col := docVisits.Collection(counter)
	// Construct a reference to the document.
	doc := col.Doc(strconv.Itoa(idx))
	// Transactionally increase the counter, or set it to one if it does not exist.
	return repo.client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// Get the current value.
		snapshot, err := doc.Get(ctx)

		if err != nil && !strings.Contains(err.Error(), "NotFound") {
			// There was an error and it was not NotFound, so we return.
			return err
		}
		if !snapshot.Exists() {
			// Create the shard and set its value to 1.
			_, err = doc.Set(ctx, docShard{1})
		} else {
			// Get the data, increase the value, and write back.
			shard := createShard(snapshot.Data())
			shard.Count++
			_, err = doc.Set(ctx, shard)
		}
		return err
	})
}

type docShard struct {
	Count int64
}

func createShard(data map[string]interface{}) docShard {
	return docShard{data["Count"].(int64)}
}
