// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
package main

import (
	"fmt"
)

func main() {
	x := 100
	fmt.Printf("%d\t %b\t %#x\n", x, x, x)
}
