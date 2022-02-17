/*
Declare two variables a and b with values 14 and 32. Swap their values such that a contains 32 and b contains 14.
*/
package main

import "fmt"

func main() {
	a := 14
	b := 32

	a, b = b, a // my favorite trick

	fmt.Println("a", a)
	fmt.Println("b", b)
}
