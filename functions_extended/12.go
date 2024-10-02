package functions_extended

import "fmt"

func Example13() {
	fmt.Println("Начало работы")
	safeFunction()
	fmt.Println("Продолжение после recover")
}

func safeFunction() {
	x := func() {
		fmt.Println("DEFER")
		if r := recover(); r != nil {
			fmt.Println("Recover from:", r)
		}
	}
	defer x()

	fmt.Println("До вызова panic")
	panic("что-то пошло не так")
	fmt.Println("Этот код не будет выполнен")
}
