package main

import (
	"fmt"
	"github.com/rhaulkumar-techo/olx/internal/config"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.MustLoad()

	// custom router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(`{"status":"running"}`))
	})

	// server config
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Println("Olx server is running on port 6000...")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
