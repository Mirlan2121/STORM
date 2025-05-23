package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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

	// Логирование запроса INFO
	log.Printf("[INFO] %s %s | Client: %s | Time: %v",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		time.Since(start))

	file, _ := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(file)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s | Duration: %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
