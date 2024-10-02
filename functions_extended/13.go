package functions_extended

import "fmt"

func Example14() {
	funcs := []func(){}
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Printf("%p\n", &i)
		})
	}

	/*
		Все функции выведут 3, поскольку замыкание захватывает одну и ту же переменную `i`
	*/
	for _, f := range funcs {
		f()
	}
}
