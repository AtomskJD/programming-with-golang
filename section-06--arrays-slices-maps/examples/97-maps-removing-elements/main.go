package main

import "fmt"

func main() {
	playerHPs := map[string]int{
		"coolBoy123": 10,
		"aldarien":   50,
		"badBoyzz":   30,
	}
	delete(playerHPs, "coolBoy123")
	fmt.Println(playerHPs)
}
