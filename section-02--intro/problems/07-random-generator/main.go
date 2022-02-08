/*
Create a go program which generates a random number
between 0 and 10 inclusively.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(11))
}
