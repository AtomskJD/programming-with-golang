package main

import "fmt"

func main() {
	cond1 := true
	cond2 := false
	cond3 := false
	cond4 := false

	if cond1 && cond2 && cond3 && cond4 {
		fmt.Println("All conditions match")
	} else {
		fmt.Println("At least one conditions don't match")
	}

	if cond1 || cond2 || cond3 || cond4 {
		fmt.Println("At least one condition match")
	} else {
		fmt.Println("Neither of the conditions match")
	}

	carType := "Automatic"
	gearsCnt := 0
	if (carType == "Automatic" && gearsCnt == 0) ||
		(carType == "Manual" && gearsCnt > 0) {
		fmt.Println("Correct configuration")
	} else {
		fmt.Println("The configuration is incorrect")
	}

	fmt.Println(true && true)
}
