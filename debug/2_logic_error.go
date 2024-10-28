package debug

import "fmt"

// Ошибка в логике: результат сложения неверен
func add(a, b int) int {
	return a - b // Должно быть a + b
}

func LogicError() {
	result := add(2, 3)
	fmt.Println("Результат:", result) // Ожидается 5, но будет -1
}
