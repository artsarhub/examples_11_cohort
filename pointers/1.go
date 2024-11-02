package pointers

import "fmt"

func Example1() {
	var a int = 10
	var p *int = &a // Указатель p указывает на переменную a

	fmt.Println("Адрес a:", &a) // Адрес a в памяти
	a = 19
	*p = 18

	fmt.Println("Значение a:", a)            // 10
	fmt.Println("Адрес a:", &a)              // Адрес a в памяти
	fmt.Println("Значение указателя p:", p)  // Адрес, на который указывает p
	fmt.Println("Значение по адресу p:", *p) // 10 (значение a)
}
