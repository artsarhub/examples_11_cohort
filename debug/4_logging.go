package debug

import (
	"log"
	"os"
)

func Logging() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()

	log.Println("Файл успешно открыт")
}
