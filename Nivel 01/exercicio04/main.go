package main

import (
	"fmt"
)

type inteiro int

var x inteiro

func main() {
	fmt.Printf("O valor é %v e o tipo é %T\n", x, x)
	x = 42
	fmt.Println(x)
}
