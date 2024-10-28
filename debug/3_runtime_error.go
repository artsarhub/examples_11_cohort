package debug

import "fmt"

func RuntimeError() {
	arr := []int{1, 2, 3}
	fmt.Println(arr[5]) // Доступ к несуществующему индексу вызовет панику
}
