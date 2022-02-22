package main

import "fmt"

func main() {
	start, end := 10, 20
	iterate(start, end)
	fmt.Println("iterate", start, end)

	iteratePtr(&start, &end)
	fmt.Println("iterate with pointer", start, end)

}

func iterate(start, end int) {
	for start < end {
		fmt.Println(start)
		start++
	}
}

func iteratePtr(start, end *int) {
	for *start < *end {
		fmt.Println(*start)
		*start++
	}
}
