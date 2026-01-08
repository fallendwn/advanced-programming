package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DenisLi/Assignment2/internal/api"
	"github.com/DenisLi/Assignment2/internal/model"
	"github.com/DenisLi/Assignment2/internal/queue"
	"github.com/DenisLi/Assignment2/internal/store"
	"github.com/DenisLi/Assignment2/internal/worker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	taskStore := store.NewStore()
	taskQueue := queue.NewQueue[*model.Task](100)

	stopWorkers := make(chan struct{})
	for i := 0; i < 2; i++ {
		go worker.Worker(i, taskStore, taskQueue, stopWorkers)
	}
	go func() {
		timer := time.NewTicker(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				submitted, completed, progress := taskStore.GetStats()
				log.Printf("statistics: submitted=%d in_progress=%d completed=%d", submitted, progress, completed)
			case <-ctx.Done():
				return
			}
		}
	}()
	apiHandler := api.NewApi(taskStore, taskQueue)

	http.HandleFunc("/tasks", apiHandler.TaskHandler)
	http.HandleFunc("/tasks/", apiHandler.GetTaskById)
	http.HandleFunc("/stats", apiHandler.GetStats)
	server := &http.Server{

		Addr: ":8000",
	}
	go func() {

		log.Println("Starting server on 8000")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {

			log.Fatalf("Server is failed to start %w", err)

		}

	}()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	log.Println("Shutting down the server...")

	close(stopWorkers)
	taskQueue.Close()
	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer shutdownCancel()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("Could not shutdown the server : %w", err)
	}
	log.Println("..............")
}
