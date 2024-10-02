package functions_extended

import (
	"fmt"
	"os"
)

func Example8() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close() // Закроется в конце функции

	// Работа с файлом
	fmt.Println("Файл открыт успешно")
	return
}
