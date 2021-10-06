package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"aba.taller3.fi.uba.ar/tp1/visitcounter/components/storage"
	"aba.taller3.fi.uba.ar/tp1/visitcounter/components/visitcount"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		fmt.Println("The program finished successfully")
	}
}

func run() error {
	var err error

	// Initialize the root context.
	rootContext := context.Background()
	// Initialize a cancellable context and the cancel function.
	ctx, cancel := context.WithCancel(rootContext)
	defer cancel()

	var repository storage.VisitCounterRepository

	// Initialize visit counter repository.
	if repository, err = initializeVisitCounterRepository(ctx); err != nil {
		return err
	}
	defer repository.Close()

	// Initialize controller.
	controller := visitcount.CreateVisitCountController(repository)

	// Initialize router.
	router := gin.Default()
	router.GET("/api/visits/total", controller.GetTotal)

	// Initialize server.
	return listen(ctx, router)
}

func listen(ctx context.Context, router *gin.Engine) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// Launch server in a separate goroutine.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}()

	// Wait for the quit signal.
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, os.Interrupt)
	<-quitChannel

	// Try to stop the server for five seconds; if it has not stopped by then just quit.
	cctx, ccancel := context.WithTimeout(ctx, 5*time.Second)
	defer ccancel()

	if err := server.Shutdown(cctx); err != nil {
		return err
	}
	return nil
}

func initializeVisitCounterRepository(ctx context.Context) (storage.VisitCounterRepository, error) {
	config := storage.FirestoreVisitCounterRepositoryConfig{
		ProjectId: os.Getenv("GCP_PROJECT_ID"),
	}
	return storage.CreateFirestoreVisitCounterRepository(ctx, &config)
}
