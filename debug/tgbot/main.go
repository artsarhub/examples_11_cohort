package main

import (
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	token := ""

	log.Infof("Telegram Bot started [%s]", token)
	// Замените TOKEN на ваш реальный токен
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	// Включаем логирование для бота
	//bot.Debug = true

	log.Infof("Авторизован под именем %s", bot.Self.UserName)

	err = func() error {
		return err
	}()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Настраиваем канал для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Обрабатываем каждое сообщение
	for update := range updates {
		if update.Message != nil { // Если есть сообщение
			log.Debugf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			// Создаем ответное сообщение
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			// Отправляем сообщение
			bot.Send(msg)
		}
	}
}
