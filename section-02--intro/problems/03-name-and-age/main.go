/*
Create a program, which prints your first and last names on the terminal,
followed by your age.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hi my name Ilya")
	t := time.Now()
	fmt.Println("my age is", t.Year()-1984)
}
