package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/trooffEE/sushi-clicker-backend/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*10, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/login", handlers.Login).Methods(http.MethodPost, http.MethodOptions)

	// CORS middleware
	router.Use(mux.CORSMethodMiddleware(router))

	// TODO hide it, too much logic in one place

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3010",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
