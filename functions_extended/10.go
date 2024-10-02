package functions_extended

import "fmt"

func Example10() {
	defer fmt.Println("Первый отложенный вызов")
	defer fmt.Println("Второй отложенный вызов")
	defer fmt.Println("Третий отложенный вызов")

	fmt.Println("Основной код функции")
	// Output:
	// Основной код функции
	// Третий отложенный вызов
	// Второй отложенный вызов
	// Первый отложенный вызов
}
