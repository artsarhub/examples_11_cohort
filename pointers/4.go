package pointers

import "fmt"

func Example4() {
	var p *int
	fmt.Println(&p)
	if p != nil {
		fmt.Println("Указатель не нулевой:", *p)
	} else {
		fmt.Println("Указатель равен nil")
	}

	var num int = 10
	fmt.Println(&num)
	p = &num
	var z **int = &p
	fmt.Println(&z)
	fmt.Println(*z)
	fmt.Println(**z)
	var zz ***int = &z
	fmt.Println(&zz)
	fmt.Println(***zz)
	if p != nil {
		fmt.Println("Указатель не нулевой:", *p) // 10
	}
}
