package functions_extended

import "fmt"

// Функция для вычисления факториала
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func Example6() {
	fmt.Println("Факториал 5:", factorial(5)) // Output: Факториал 5: 120
}
