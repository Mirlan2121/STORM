package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"status": "I AM THE STORM"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Настройка лимитера: 100 запросов в минуту
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}
	store := memory.NewStore()
	middleware := stdlib.NewMiddleware(limiter.New(store, rate))

	http.Handle("/", middleware.Handler(http.HandlerFunc(handleRequest)))
	log.Println("Server started on :8080")
	http.ListenAndServeTLS(":443", "tls/cert.pem", "tls/key.pem", nil)
	// Меняем порт на 8443 (не требует прав root)
	http.ListenAndServeTLS(":8443", "tls/cert.pem", "tls/key.pem", nil)

	// Запуск HTTPS
	log.Println("Starting HTTPS server on :443")
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("HTTPS failed: ", err)
	}
}
