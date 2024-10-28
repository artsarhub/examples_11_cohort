package animals

import (
	"fmt"

	"yp-examples/interfaces/logger"
)

type cat struct {
	AnimalStruct
}

func NewCat(name string, fullness int, log logger.Logger) cat {
	cat := cat{}
	cat.name = name
	cat.fullness = fullness
	cat.log = log
	return cat
}

func (c *cat) Feed() {
	c.fullness++
	c.log.Info(fmt.Sprintf("So yummy! My fullness is %d.", c.fullness))
}

func (c *cat) SayHello() {
	c.log.Info(fmt.Sprintf("Hi! My name is %s. I'm cat.", c.name))
	c.fullness -= 2
}
