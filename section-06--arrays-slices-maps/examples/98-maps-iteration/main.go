package main

import "fmt"

func main() {
	primes := map[int]bool{
		2: true,
		3: false,
		4: false,
		5: true,
	}

	for key, value := range primes {
		fmt.Println(key, value) // order not guaranteed
	}
}
