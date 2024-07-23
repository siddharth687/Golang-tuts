package main

import "fmt"

func main() {
	fmt.Println("If Else in Golang")

	loginCount := 24
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "New user"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

}
