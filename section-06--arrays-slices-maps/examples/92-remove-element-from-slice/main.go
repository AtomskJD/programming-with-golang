package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Katya", "George", "Lucy"}
	toRemove := 2 // remove George

	zombiesWithoutGeorge := make([]string, len(zombies[:toRemove]))
	copy(zombiesWithoutGeorge, zombies[:toRemove])

	zombiesWithoutGeorge = append(zombiesWithoutGeorge[:toRemove], zombies[toRemove+1:]...)

	fmt.Println(zombies)
	fmt.Println(zombiesWithoutGeorge)

}
