package main

import "fmt"

func main() {
	printPlayerName("test", "Lowko", "Nemo")
}

func printPlayerName(players ...string) {
	for _, player := range players {
		fmt.Printf("Player %s is cool!\n", player)
	}
}
