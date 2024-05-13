package main

import (
	"context"
	"errors"
	"github.com/lfordyce/generalNotes/cmd"
	"github.com/lfordyce/generalNotes/cmd/service/monster"
	"log"
	"net/http"
)

func main() {
	ac := make(chan server, 1)
	setupServer(ac)
	ps := <-ac

	ctx, cancel := context.WithCancel(context.Background())
	if err := handleServer(ps)(ctx, cancel); err != nil {
		panic(err)
	}
}

func setupServer(ac chan server) {
	handler := new(monster.Handler)

	//router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Received a non domain request")
	//	if _, err := w.Write([]byte("Hello, stranger...")); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	}
	//}))
	router := http.NewServeMux()
	router.HandleFunc("POST /monster", handler.Create)
	router.HandleFunc("PUT /monster/{id}", handler.UpdateByID)
	router.HandleFunc("GET /monster/{id}", handler.FindByID)
	router.HandleFunc("DELETE /monster/{id}", handler.DeleteByID)

	svc := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	ac <- svc
}

type Start func(ctx context.Context, cancel context.CancelFunc) error

func handleServer(s server) Start {
	return func(ctx context.Context, cancel context.CancelFunc) error {
		exitCh := make(chan error)

		go cmd.WaitForTerminate(ctx, func(msg string) {
			log.Printf("terminate signal: %s\n", msg)
			exitCh <- s.Shutdown(ctx)
		})

		defer cancel()
		serverErr := s.ListenAndServe()

		// When Shutdown is called, Serve, ListenAndServe, and
		// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
		// program doesn't exit and waits instead for Shutdown to return.
		if errors.Is(serverErr, http.ErrServerClosed) {
			// waits instead for Shutdown to return.
			if err := <-exitCh; err != nil {
				log.Printf("shutdown error: %v\n", err)
			}
			// on shutdown, no error or error already logged
			return nil
		}
		return serverErr
	}
}
