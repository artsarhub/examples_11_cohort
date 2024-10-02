package functions_extended

import "fmt"

func Example4() {
	// Создание замыкания
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}

	var z bool
	fmt.Println(z)

	count := counter()
	fmt.Println(count()) // Output: 1
	fmt.Println(count()) // Output: 2
	fmt.Println(count()) // Output: 3

	count = func() int {
		fmt.Println("HI")
		return 1
	}
	fmt.Println(count()) // Output: 1
	fmt.Println(count()) // Output: 2
	fmt.Println(count()) // Output: 3

	//count2 := counter()
	//fmt.Println(count2())
	//fmt.Println(count2())
}
