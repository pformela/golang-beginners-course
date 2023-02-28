package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type Animal2 struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	b2 := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(b2)

	t := reflect.TypeOf(Animal2{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
