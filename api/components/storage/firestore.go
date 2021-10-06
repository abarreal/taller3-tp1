package storage

import (
	"context"
	"errors"
	"strings"

	"cloud.google.com/go/firestore"
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

func (repo *firestoreVisitCounterRepository) GetCount(ctx context.Context) (uint64, error) {
	// The total count is stored in the document /counters/aggregates, in field "visits"
	col := repo.client.Collection("counters")
	doc := col.Doc("aggregates")

	snapshot, err := doc.Get(ctx)

	if err != nil && !strings.Contains(err.Error(), "NotFound") {
		return 0, err
	} else if !snapshot.Exists() {
		// The document does not exist, so there are no visits yet.
		return 0, nil
	}
	// Get the data.
	if visits, found := snapshot.Data()["visits"]; !found {
		return 0, errors.New("unexpected document format")
	} else {
		return uint64(visits.(int64)), nil
	}
}
