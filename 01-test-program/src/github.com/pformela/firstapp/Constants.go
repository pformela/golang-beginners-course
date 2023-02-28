package main

import (
	"fmt"
)

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	specialistType = catSpecialist
	fmt.Printf("%v, %T\n", specialistType, specialistType)
}
