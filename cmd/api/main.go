package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rhaulkumar-techo/olx/internal/config"
	"github.com/rhaulkumar-techo/olx/internal/handlers"
)

func main() {
	cfg := config.MustLoad()

	// custom router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /",handlers.Healthz)

	// server config
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Printf("Olx server is running on port %s...\n", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
