package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Katya", "George", "Lucy"}
	toRemove := 2 // remove George

	// just short write without any variables
	result := append(zombies[:toRemove], zombies[toRemove+1:]...)

	// write Lucy to idx 2 original slice
	// zombies = [Paul Katya Lucy Lucy] result = [Paul Katya Lucy]
	// более безопасно присваивать оригинальному слайсу
	// иначе оригинальный слайс будет неожиданным
	fmt.Println(zombies, result)
}
