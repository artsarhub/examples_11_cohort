package functions_extended

import "fmt"

func Example11() {
	x := 10
	defer func() {
		fmt.Println("Значение x при отложенном вызове:", x)
	}()

	x += 5
	fmt.Println("Значение x после изменения:", x)
	// Output:
	// Значение x после изменения: 15
	// Значение x при отложенном вызове: 10
}
