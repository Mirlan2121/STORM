package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// Добавляем SecurityMiddleware без изменения существующих функций
func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Базовые security headers
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		next.ServeHTTP(w, r)
	})
}

// Оригинальный handleRequest без изменений
func handleRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response := map[string]string{"status": "I AM THE STORM"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	log.Printf(
		"Request: %s %s | From: %s | Duration: %v",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
		time.Since(start),
	)

	file, _ := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)
}

// Оригинальный LoggingMiddleware без изменений
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request: %s %s | Duration: %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	// Обертываем цепочку middleware с сохранением оригинальной логики
	handler := SecurityMiddleware(
		LoggingMiddleware(
			http.HandlerFunc(handleRequest),
		),
	)

	http.Handle("/", handler)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
