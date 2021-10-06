package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"aba.taller3.fi.uba.ar/tp1/visitcounter/components/storage"
	"cloud.google.com/go/logging"
)

const LoggerName = "aggregator"

func main() {

	// Initialize the root context.
	ctx := context.Background()

	// Initialize Cloud logger.
	loggingClient, err := logging.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"))

	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not initialize logger")
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer loggingClient.Close()
	logger := loggingClient.Logger(LoggerName)

	if err := run(ctx, logger); err != nil {
		errlog(logger, err)
		os.Exit(1)
	} else {
		fmt.Println("The program finished successfully")
	}
}

func run(rootContext context.Context, logger *logging.Logger) error {
	var err error

	// Initialize a cancellable context and the cancel function.
	ctx, cancel := context.WithCancel(rootContext)
	defer cancel()

	// Initialize visit counter repository.
	log(logger, "Initializing visit counter repository")
	var repository storage.VisitCounterRepository

	if repository, err = initializeVisitCounterRepository(ctx); err != nil {
		return err
	}
	defer repository.Close()

	// Initialize aggregation configuration.
	aggregationInterval, err := strconv.Atoi(os.Getenv("AGGREGATION_INTERVAL_IN_SECONDS"))
	if err != nil {
		return err
	}

	// Listen for quit signals.
	log(logger, "Initializing signal handler")
	go handleSignals(cancel, logger)

	// Repeat the aggregation procedure every N seconds until a quit
	// signal is received.
	stopping := false

	log(logger, fmt.Sprintf("Aggregation interval: %d seconds", aggregationInterval))
	timer := time.NewTicker(time.Duration(aggregationInterval) * time.Second)
	defer timer.Stop()

	log(logger, "Starting main loop")
	for !stopping {

		log(logger, "Sleeping until new aggregation cycle")

		// GCP container VM seems to have some issue with the following code.
		// The container simply dies silently without executing either
		// logging statement.
		/*
			select {
			case <-ctx.Done():
				log(logger, "The program is stopping")
				stopping = true
				continue
			case <-timer.C:
				// Perform aggregation procedures.
				log(logger, "Running aggregation procedures")
				if err := repository.AggregateCounters(ctx); err != nil {
					return err
				}
			}
		*/

		// Wait for a new timer signal.
		<-timer.C

		log(logger, "Running aggregation procedures")
		if err := repository.AggregateCounters(ctx); err != nil {
			return err
		}
	}

	log(logger, "Main loop finished")
	return nil
}

func handleSignals(cancel context.CancelFunc, logger *logging.Logger) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	<-channel
	log(logger, "Stop signal received")
	cancel()
}

func initializeVisitCounterRepository(ctx context.Context) (storage.VisitCounterRepository, error) {
	config := storage.FirestoreVisitCounterRepositoryConfig{
		ProjectId: os.Getenv("GCP_PROJECT_ID"),
	}
	return storage.CreateFirestoreVisitCounterRepository(ctx, &config)
}

func log(logger *logging.Logger, msg string) {
	logger.Log(logging.Entry{
		Payload: msg,
	})
}

func errlog(logger *logging.Logger, err error) {
	logger.Log(logging.Entry{
		Payload:  err.Error(),
		Severity: logging.ParseSeverity("Error"),
	})
}
