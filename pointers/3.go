package pointers

import "fmt"

type Person struct {
	name string
	age  int
}

func changeName(p *Person, newName string) {
	p.name = newName
}

func Example3() {
	person := Person{name: "Alice", age: 25}
	fmt.Println("Имя до изменения:", person.name) // Alice

	changeName(&person, "Bob")
	fmt.Println("Имя после изменения:", person.name) // Bob
}
