package animals

import (
	"fmt"

	"yp-examples/interfaces/logger"
)

type Dog struct {
	AnimalStruct
}

func NewDog(name string, fullness int, log logger.Logger) Dog {
	dog := Dog{}
	dog.name = name
	dog.fullness = fullness
	dog.log = log
	return dog
}

func (d Dog) Feed() {
	d.fullness++
	d.log.Error(fmt.Sprintf("So yummy! My fullness is %d.", d.fullness))
}

func (d Dog) SayHello() {
	d.log.Warn(fmt.Sprintf("Hi! My name is %s. I'm dog.", d.name))
	d.fullness -= 5
}
