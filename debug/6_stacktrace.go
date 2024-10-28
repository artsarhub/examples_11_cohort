package debug

import (
	"fmt"
)

func StacktraceReading() {
	fmt.Println("Пример многомодульной программы")
	res := Divide(4, 0)
	fmt.Println(res) // Вызовет деление на ноль
}

func Divide(a, b int) int {
	return a / b // Здесь возникнет ошибка времени выполнения
}
