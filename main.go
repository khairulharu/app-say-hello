package main

import (
	"fmt"

	gosayhello "github.com/khairulharu/go-say-hello/v3"
)

func main() {
	var hallo string = gosayhello.SayHello()
	fmt.Println(hallo)

	fmt.Println(gosayhello.SayNumber(6))

	fmt.Println(gosayhello.SayMultiple(2))
}
