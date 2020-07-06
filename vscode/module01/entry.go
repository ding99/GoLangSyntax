package main

import (
	"fmt"

	"wtest.com/modules/module01/utilities"
)

func main() {
	fmt.Println("-- Entrance : Hello")

	fmt.Println(utilities.ReverseRunes("!tsetw olleH"))
	utilities.OddNumbers()
}
