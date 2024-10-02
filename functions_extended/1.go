package functions_extended

import "fmt"

func Example1() {
	// Анонимная функция, присвоенная переменной
	add := func(a, b int) int {
		return a + b
	}

	result := add(5, 3)
	fmt.Println("Результат сложения:", result) // Output: Результат сложения: 8

	// Вызов анонимной функции прямо при её объявлении
	func() {
		fmt.Println("Hello from an anonymous function")
	}()
}
