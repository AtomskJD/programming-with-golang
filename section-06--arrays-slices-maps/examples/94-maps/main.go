package main

import "fmt"

func main() {
	playerHPs := map[string]int{
		"coolBoy123": 10,
		"aldarien":   50,
		"badBoyzz":   30,
	}
	hp := playerHPs["badBoyzz"]
	fmt.Println(hp)
	playerHPs["badBoyzz"] = 25
	fmt.Println(playerHPs["badBoyzz"])

	hp, ok := playerHPs["badBoyzz"]
	if !ok {
		panic("element is not exists") // if not exist
	}
	fmt.Println(hp)

	//var nonInitMap map[string]int
	// nonInitMap["someKey"] = 42 // get an error
	// but
	initMap := map[string]int{}
	initMap["someKey"] = 42 // don't get an error

	_, ok = initMap["someNewKey"]
	if !ok {
		initMap["someNewKey"] = 32
	}

	if _, ok = initMap["someNewNewKey"]; !ok {
		initMap["someNewNewKey"] = 11
	}

	fmt.Println(initMap)

}
