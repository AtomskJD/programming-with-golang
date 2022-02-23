package main

import (
	"fmt"
	"time"
)

type Person struct {
	name      string
	age       int
	dateBirth time.Time
}

func main() {
	var p Person
	p.name = "Pepe"
	p.age = 25
	p.dateBirth = time.Now()

	p2 := Person{
		name:      "Dodo",
		age:       33,
		dateBirth: time.Now(),
	}

	p3 := p
	p3.age = 100 // this time not a pointer

	fmt.Println(p, p2, p3)
}
