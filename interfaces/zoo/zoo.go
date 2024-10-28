package zoo

import (
	"fmt"

	"yp-examples/interfaces/logger"
	"yp-examples/interfaces/zoo/animals"
)

//go:generate mockgen -package=mock -destination=./mock/animal_mock.go yp-examples/interfaces/zoo Animal
type Animal interface {
	Feed()
	SayHello()
	GetFullness() int
}

func EmulateZoo(log logger.Logger) {
	var animalsSlice []Animal

	for _, name := range []string{"Bars", "Murzik"} {
		newCat := animals.NewCat(name, 100, log)
		animalsSlice = append(animalsSlice, &newCat)
	}
	for _, name := range []string{"Sharik", "Rex"} {
		newDog := animals.NewDog(name, 100, log)
		animalsSlice = append(animalsSlice, &newDog)
	}

	for _, a := range animalsSlice {
		a.SayHello()
		log.Info(fmt.Sprintf("My fullness is %d for now.", a.GetFullness()))
		a.Feed()
	}
}
