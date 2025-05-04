package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HealisticEngineer/CAFE/db"
	"github.com/HealisticEngineer/CAFE/handlers"
)

func main() {
	db.InitDB()
	db.EnsureTables() // Ensure tables exist before starting the server

	// Add health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Register the config handler
	http.HandleFunc("/config", handlers.ConfigHandler)

	// Return a 404 for any other route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	srv := &http.Server{
		Addr:              ":3000",
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Configuration server is running on port 3000")
	log.Fatal(srv.ListenAndServe())
}
