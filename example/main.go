package main

import (
	"fmt"

	"github.com/jojomi/ringcounter"
)

func main() {
	c := ringcounter.New(7, 0)

	for i := 0; i < 100; i++ {
		fmt.Println("Round:", i, ", Counter: ", c.Value)
		c.Advance()
	}
}
