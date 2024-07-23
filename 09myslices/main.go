package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video on slices")
	var fruitlist = []string{"Apple", "Banana", "Peach"}
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[1:])

	fmt.Println(fruitlist)
}
