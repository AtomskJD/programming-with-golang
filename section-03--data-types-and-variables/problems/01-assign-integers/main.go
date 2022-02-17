/*
Create a program, which assigns a set of integer
values to the appropriate type in Go. Print the
integer on the terminal. The numbers should be
stored in the smallest type possible.

*/
package main

import "fmt"

func main() {
	var i1 uint8 = 60
	fmt.Println("uint8 ", i1)

	var i2 int8 = -100
	fmt.Println("int8 ", i2)

	var i3 int8 = 127
	fmt.Println("int8 ", i3)

	var i4 uint8 = 128
	fmt.Println("uint8", i4)

	var i5 int32 = -144243
	fmt.Println("int32", i5)

	var i6 uint8 = 255
	fmt.Println("uint8", i6)

	var i7 int16 = 256
	fmt.Println("int16", i7)

	var i8 int32 = 144243
	fmt.Println("int32", i8)

	var i9 int16 = 3641
	fmt.Println("int16", i9)

	var i10 int16 = -4512
	fmt.Println("int16", i10)

	var i11 uint16 = 65535
	fmt.Println("uint16", i11)

	var i12 uint64 = 10000000000000000000
	fmt.Println("int32", i12)

	var i13 int32 = 65536
	fmt.Println("int32", i13)

	var i14 int64 = -1000000000000000000
	fmt.Println("int64", i14)

	var i15 int16 = -30000
	fmt.Println("int16", i15)
}
