package animals

import (
	"fmt"

	"yp-examples/interfaces/logger"
)

type AnimalStruct struct {
	name     string
	fullness int
	log      logger.Logger
}

func (c *AnimalStruct) GetFullness() int {
	return c.fullness
}

func (c *AnimalStruct) SetFullness(v int) {
	if v < 0 {
		fmt.Println("SetFullness called with negative value")
		return
	}
	c.fullness = v
}
