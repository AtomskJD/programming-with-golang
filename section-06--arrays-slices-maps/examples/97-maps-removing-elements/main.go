package main

import "fmt"

func main() {
	playerHPs := map[string]int{
		"coolBoy123": 10,
		"aldarien":   50,
		"badBoyzz":   30,
	}
	playerHPs2 := playerHPs
	playerHPs2["aldarien"] = 17
	playerHPs2["test"] = 33
	delete(playerHPs2, "coolBoy123")
	fmt.Println(playerHPs, playerHPs2) // опять поинтер
}
