package main

import "fmt"

func main() {
	mp := map[string]int{"foo": 1}
	updateMap(mp)
	fmt.Println(mp)

	sl := []int{1, 2, 3, 4, 5}
	updateSlice(sl)
	fmt.Println(sl)
}

func updateMap(mp map[string]int) {
	mp["foo"] = 11
	mp["bar"] = 22

}

func updateSlice(sl []int) {
	for i := range sl {
		sl[i] += 1
	}
}
