package main

import "fmt"

func main() {
	x := []int{1, 2}
	y := []int{3, 4}
	ref := x
	fmt.Println(x, y, ref) // [1 2] [3 4] [1 2]
	x = y                  // x теперь тотже указатель что и y а старый х указатель есть только у ref
	x[0] = 5
	fmt.Println(x, y, ref) // [5 4] [5 4] [1 2]
}
