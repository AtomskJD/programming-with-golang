/*
Convert the following values to the type specified & print them out:
*/
package main

import "fmt"

func main() {
	var fl1 float32 = 3.14
	fmt.Println(int(fl1))

	var fl2 float64 = 12.3456789
	fmt.Println(float32(fl2))

	var i1 int16 = 1234
	fmt.Println(int8(i1))

	var i2 int16 = 1234
	fmt.Println(uint8(i2))
}
