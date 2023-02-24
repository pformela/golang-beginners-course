package main

import (
	"fmt"
)

func main() {
	var n bool = false

	m := 1 == 1
	k := 1 == 2

	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", k, k)

	fmt.Printf("%v, %T\n", n, n)

	var j uint16 = 32
	fmt.Printf("%v, %T\n", j, j)

	s := "this is a string"
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

}
