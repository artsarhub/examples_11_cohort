package functions_extended

import "fmt"

// Функция, создающая другую функцию для умножения на заданное число
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func Example7() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println("Удвоение 5:", double(5)) // Output: Удвоение 5: 10
	fmt.Println("Утроение 4:", triple(4)) // Output: Утроение 4: 12
}
