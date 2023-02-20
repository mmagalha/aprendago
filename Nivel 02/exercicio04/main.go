package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Printf("%d\t %b\t %#x\n", x, x, x)
	y := x << 2
	fmt.Printf("%d\t %b\t %#x\n", y, y, y)
}
