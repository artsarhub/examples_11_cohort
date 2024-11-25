package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"yp-examples/interfaces/logger"
)

const (
	protocol = "http"
	host     = "localhost"
	port     = 8081
)

var log logger.Logger

// Message структура для данных, которые будут отправляться в POST-запросе и получаться в ответе
type Message struct {
	Text string `json:"text"`
	User string `json:"user"`
}

func sendGetRequest(url string) {
	// Отправляем GET-запрос к серверу
	resp, err := http.Get(url)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка GET-запроса: %s", err))
		return
	}
	defer resp.Body.Close()

	// Читаем и выводим ответ от сервера на GET-запрос
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка чтения ответа на GET-запрос: %s", err))
		return
	}
	log.Info(fmt.Sprintf("Ответ от сервера (GET): %s", string(body)))
}

func sendPostRequest(url string) {
	// Создаем данные для отправки в формате JSON
	message := Message{
		Text: "Привет, сервер!",
		User: "Клиент 1",
	}

	// Сериализация структуры в JSON для отправки в POST-запросе
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка сериализации JSON: %s", err))
		return
	}

	// Отправляем POST-запрос с JSON-данными
	resp, err := http.Post(fmt.Sprintf("%s/post", url), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка POST-запроса: %s", err))
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Error(err.Error())
			return
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка чтения ответа: %s", err))
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Error(fmt.Sprintf("Код ответа не ОК: %d Сообщение: %s", resp.StatusCode, body))
		return
	}

	// Читаем и десериализуем JSON-ответ от сервера
	var response Message
	if err := json.Unmarshal(body, &response); err != nil {
		log.Error(fmt.Sprintf("Ошибка десериализации ответа JSON: %s Ответ: %s", err, body))
		return
	}

	log.Info(fmt.Sprintf("Ответ от сервера (POST): Пользователь=%s, Сообщение=%s\n", response.User, response.Text))
}

func main() {
	log = logger.NewConsoleLogger()
	url := fmt.Sprintf("%s://%s:%d", protocol, host, port)

	sendGetRequest(url)
	sendPostRequest(url)
}
