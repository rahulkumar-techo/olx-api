package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rhaulkumar-techo/olx/internal/handlers"
	"github.com/rhaulkumar-techo/olx/internal/platform/config"
	"github.com/rhaulkumar-techo/olx/internal/platform/database"
)

func main() {
	cfg := config.MustLoad()

	_, error := database.Connect(cfg.DatabaseUrl)
	if error != nil {
		log.Fatalf("main.db.connect: %v", error)
	}
	fmt.Println("DATABASE connected")

	// custom router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Healthz)

	// server config
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Olx server is running on port %s...", cfg.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
