package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Simulate a slow startup by sleeping for 30 seconds
	fmt.Println("Starting application...")
	time.Sleep(60 * time.Second)
	fmt.Println("Application started!")

	// Define the health check handler for "/healthz"
	http.HandleFunc("/application/health", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a 200 OK to indicate that the app is healthy
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start the HTTP server on port 8888
	fmt.Println("Listening on port 8888...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}

