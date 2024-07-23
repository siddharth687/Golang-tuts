package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	Print()
	res := Adder(2, 4)
	fmt.Println("Addion of 2 numbers is", res)
	res2, mymessage := Adders(1, 2, 3, 4, 5, 6)
	fmt.Println("Addition of multiple numbers is", res2)
	fmt.Println(mymessage)
}

func Adders(v ...int) (int, string) {
	var sum int
	// for i := 0; i < len(v); i++ {
	// 	sum += v[i]
	// }
	for _, val := range v {
		sum += val
	}
	return sum, "Hi bro"
}

func Adder(a int, b int) int {
	return a + b
}

func Print() {
	fmt.Println("Hello to the functions")
}
