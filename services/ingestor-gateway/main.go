package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ingest", ingestHandler)

	fmt.Println("Starting ingestor-gateway on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func ingestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Placeholder: In the future, this will validate, enrich, and publish the event.
	log.Println("Received an event.")

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Event accepted")
}
