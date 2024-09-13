package functions

import "fmt"

func RunAnonymousFunction() {
	//anon := func() {
	//	fmt.Println("Hello from anonymous function with name")
	//}
	//anon()
	//
	//func() {
	//	fmt.Println("Hello from anonymous function")
	//}()

	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println()
	fmt.Println()

	counter2 := createCounter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
}

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
