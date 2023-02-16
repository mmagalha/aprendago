package main

import (
	"fmt"
)

type inteiro int

var x inteiro

var y int

func main() {
	fmt.Printf("O valor de x é %v e o tipo é %T\n", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("O valor de y é %v e o tipo é %T\n", y, y)
}
