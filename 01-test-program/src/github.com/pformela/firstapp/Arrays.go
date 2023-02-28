package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93}
	grades2 := [...]int{97, 85, 93}

	var students [3]string

	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades: %v\n", grades2)
	fmt.Printf("Students: %v\n", students)

	students[0] = "Lisa"

	fmt.Printf("Students: %v\n", students)

	var identityMatrix [3][3]int = [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, 1, 0},
		[3]int{0, 0, 1},
	}

	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]
	e := c[3:]
	f := c[:6]
	g := c[3:6]

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Println(len(h))
	fmt.Println(cap(h))

	h = append(h, 1)

	fmt.Println(h)
	fmt.Println(len(h))
	fmt.Println(cap(h))
}
