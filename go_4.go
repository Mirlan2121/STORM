package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now() // Засекаем время

	response := map[string]string{"status": "I AM THE STORM"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Логируем запрос
	log.Printf(
		"Request: %s %s | From: %s | Duration: %v",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		time.Since(start),
	)
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
