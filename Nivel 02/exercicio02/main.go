package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20

	r := a == b
	s := a != b
	t := a <= b
	u := a < b
	v := a >= b
	x := a > b

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", r, s, t, u, v, x)

}
