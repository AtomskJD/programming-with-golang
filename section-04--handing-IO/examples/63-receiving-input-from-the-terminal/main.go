package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)
	fmt.Println("line 1: ", line)

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.TrimSpace(line)
	fmt.Println("Line 2: ", line)

	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// when convert string - trim \n
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Println("My number is", num)

	// read float
	line, err = r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	floatNum, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("This is float: ", floatNum)

}
