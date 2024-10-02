package functions_extended

import "fmt"

// Возвращает функцию, которая добавляет заданное значение
func createAdder(value int) func(int) int {
	//fmt.Printf("str 7: value = %d\n", value)
	f := func(x int) int {
		//fmt.Printf("str 9: x = %d, value = %d\n", x, value)
		return x + value
	}
	//fmt.Printf("str 12: f = %+v\n", f)
	return f
}

func Example3() {
	var value int
	_, err := fmt.Scanln(&value)
	if err != nil {
		panic(err)
	}
	addFive := createAdder(value)
	fmt.Printf("addFive = %+v\n", addFive)
	addFive = createAdder(5)
	fmt.Printf("addFive = %+v\n", addFive)
	v := addFive(10)
	fmt.Printf("Добавление %d к 10: %d", value, v) // Output: Добавление 5 к 10: 15
}
