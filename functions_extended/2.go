package functions_extended

import "fmt"

// Функция, которая принимает другую функцию как аргумент
func applyOperation(a, b int, op func(int, int) int) int {
	fmt.Println(a)
	fmt.Println(b)
	return op(a, b)
}

func Example2() {
	// Обычная функция сложения
	add := func(a, b int) int {
		return a + b
	}
	//aaaa := func(a, b int) int {
	//	return a * b
	//}

	result := applyOperation(5, 3, add)
	//result2 := applyOperation(5, 3, aaaa)
	fmt.Println("Результат операции:", result) // Output: Результат операции: 8
}
