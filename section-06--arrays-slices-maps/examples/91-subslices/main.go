package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	subslice := numbers[len(numbers)-2:]
	fmt.Println(subslice)   // 7 8
	subslice = numbers[0:2] // 1 2
	fmt.Println(subslice)
	subslice = numbers[0:1]
	fmt.Println(subslice) // 1

	subslice[0] = 777
	subslice = append(subslice, 666)
	fmt.Println(numbers, subslice)

}
