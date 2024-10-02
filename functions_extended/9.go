package functions_extended

import "fmt"

func logStartAndEnd(name string) func() {
	fmt.Printf("Начало функции %s\n", name)
	return func() {
		fmt.Printf("Конец функции %s\n", name)
	}
}

func Example9() {
	d := logStartAndEnd("main") // Логирование выполнения функции main
	defer d()

	fmt.Println("Основная работа функции")
	// Output:
	// Начало функции main
	// Основная работа функции
	// Конец функции main
}
