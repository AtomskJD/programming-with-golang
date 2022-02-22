package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 123456
	sum, err := calcSumOfDigits(num)
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)

	sum2, err := calcSumOfDigits(-1000)
	if err != nil {
		panic(err)
	}
	fmt.Println(sum2)
}

func calcSumOfDigits(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Negative values are not supported")
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum, nil
}
