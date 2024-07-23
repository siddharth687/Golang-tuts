package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Pointers")
	// var ptr *int
	// fmt.Println(ptr)

	mynumber := 37
	var ptr = &mynumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)
	*ptr = *ptr * 2
	fmt.Println("New Value is", mynumber)

}
