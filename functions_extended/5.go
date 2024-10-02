package functions_extended

import "fmt"

// Вариативная функция, которая принимает любое количество аргументов
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Example5() {
	fmt.Println("Сумма чисел:", sum(1, 2, 3, 4, 5)) // Output: Сумма чисел: 15
}
