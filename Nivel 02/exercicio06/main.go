package main

import (
	"fmt"
)

func main() {
	const (
		_ = iota + 2023
		b
		c
		d
		e
	)
	fmt.Printf("%v\n%v\n%v\n%v\n", b, c, d, e)
}
