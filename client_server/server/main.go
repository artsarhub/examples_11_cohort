package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"yp-examples/interfaces/logger"
)

const port = 8081

var log logger.Logger

// Message структура для данных, которые будут получены в POST-запросе
type Message struct {
	Text string `json:"text"`
	User string `json:"user"`
}

// Обработчик для GET-запросов
func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(fmt.Sprintf("Принят %s запрос от %s", r.Method, r.RemoteAddr))
	_, err := w.Write([]byte("Hello, client! This is the server response for GET request."))
	if err != nil {
		log.Error(err.Error())
		return
	}
}

// Обработчик для POST-запросов с десериализацией JSON
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение и десериализация JSON-данных из тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusInternalServerError)
		return
	}

	var msg Message
	if err := json.Unmarshal(body, &msg); err != nil {
		http.Error(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	log.Info(fmt.Sprintf("Получен POST-запрос: Пользователь=%s, Сообщение=%s\n", msg.User, msg.Text))

	// Формирование ответа в виде JSON
	response := Message{
		Text: "Сообщение получено",
		User: msg.User,
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Ошибка сериализации ответа", http.StatusInternalServerError)
	}
}

func main() {
	log = logger.NewConsoleLogger()
	// Назначаем обработчики для маршрутов
	http.HandleFunc("/", helloHandler)    // Обработчик для GET-запросов
	http.HandleFunc("/post", postHandler) // Обработчик для POST-запросов

	log.Info(fmt.Sprintf("Сервер запущен на порту :%d", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Error(fmt.Sprintf("Ошибка запуска сервера: %s", err))
	}
}
