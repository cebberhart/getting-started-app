package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type response struct {
	Language string `json:"language"`
	Message  string `json:"message"`
	Status   string `json:"status"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{
		Language: "Go (net/http)",
		Message:  "Hello from a containerized Go app!",
		Status:   "running",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Go server listening on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
