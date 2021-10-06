package visitcounter

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"aba.taller3.fi.uba.ar/tp1/visitcounter/components/storage"
)

var repository storage.VisitCounterRepository

func init() {
	var err error
	if repository, err = initializeVisitCounterRepository(context.Background()); err != nil {
		log.Fatalln(err)
	}
}

func CloudFunctionsRun(ctx context.Context, msg pubSubMessage) error {
	m := incomingVisitMessage{}
	d := msg.Data

	if err := json.Unmarshal(d, &m); err != nil {
		return err
	}

	// Generate the name of the counter from the path.
	var counter string
	counter = strings.Trim(m.Path, "/")
	counter = strings.ReplaceAll(counter, "/", "-")
	// Increase the counter in question.
	if err := repository.Increase(ctx, counter); err != nil {
		return err
	}
	return nil
}

func initializeVisitCounterRepository(ctx context.Context) (storage.VisitCounterRepository, error) {
	var err error
	var shardsPerCounter int

	if shardsPerCounter, err = strconv.Atoi(os.Getenv("SHARDS_PER_COUNTER")); err != nil {
		return nil, err
	}

	config := storage.FirestoreVisitCounterRepositoryConfig{
		ProjectId:        os.Getenv("GCP_PROJECT_ID"),
		ShardsPerCounter: uint32(shardsPerCounter),
	}
	return storage.CreateFirestoreVisitCounterRepository(ctx, &config)
}

type pubSubMessage struct {
	Data []byte `json:"data"`
}

type incomingVisitMessage struct {
	Path string `json:"path"`
}
