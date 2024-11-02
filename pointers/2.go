package pointers

import "fmt"

func changeValue(num *int) {
	*num = *num * 2 // Изменяем значение через указатель
}

func Example2() {
	value := 10
	fmt.Println("До изменения:", value) // 10

	changeValue(&value)
	fmt.Println("После изменения:", value) // 20
}
