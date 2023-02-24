package main

import (
	"fmt"
	"strconv"
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// var i int = 42
	// j := 43
	// fmt.Println(i)
	// fmt.Printf("%v, %T\n", j, j)
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)

}
