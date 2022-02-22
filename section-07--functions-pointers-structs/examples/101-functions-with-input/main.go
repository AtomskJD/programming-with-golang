package main

import "fmt"

func main() {
	player1 := "Asmongold"
	printPlayerMessage(player1, "Lowko")
}

func printPlayerMessage(player1, player2 string) {
	fmt.Printf("Player %s and %s is on the starting line...\n", player1, player2)
	fmt.Println("The battle is about to begin!")
}
