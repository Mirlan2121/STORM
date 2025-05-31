package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHandleRequest(t *testing.T) {
	// Создаем тестовый запрос
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleRequest)

	// Выполняем запрос
	handler.ServeHTTP(rr, req)

	// Проверяем статус код
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем содержимое ответа
	expected := `{"status":"I AM THE STORM"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Проверяем заголовок Content-Type
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}

func TestLogFileCreation(t *testing.T) {
	// Проверяем, что файл логов создается
	if _, err := os.Stat("server.log"); os.IsNotExist(err) {
		t.Error("log file was not created")
	}
}

func TestLogging(t *testing.T) {
	// Перехватываем вывод логов
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handleRequest(rr, req)

	if !strings.Contains(buf.String(), "Request: GET /") {
		t.Errorf("Log output is wrong: %s", buf.String())
	}
}
