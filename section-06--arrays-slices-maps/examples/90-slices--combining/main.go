package main

import "fmt"

func main() {
	odds := []int{1, 3, 5, 7, 9}
	evens := []int{2, 4, 6, 8}
	result := append(odds, evens...)
	fmt.Println(result)
}
