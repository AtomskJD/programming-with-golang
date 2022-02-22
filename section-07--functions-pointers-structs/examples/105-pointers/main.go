package main

import "fmt"

func main() {
	num := 10
	fmt.Println(num)

	numCopy := num
	fmt.Println(numCopy)

	numCopy = 15
	fmt.Println(num, numCopy)

	var numPoint *int
	numPoint = &num // take address of num

	*numPoint = 77
	fmt.Println(num, numCopy, *numPoint)
}
