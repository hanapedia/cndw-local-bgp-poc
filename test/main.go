package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read POD_NAME from environment (default to "unknown" if not set)
	podName := os.Getenv("POD_NAME")
	if podName == "" {
		podName = "unknown"
	}

	// Define handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("Hello from %s!\n", podName)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(message))
	})

	addr := ":8080"
	log.Printf("Starting server on %s (pod: %s)\n", addr, podName)

	// Start HTTP server
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
