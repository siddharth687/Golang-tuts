package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file- LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylocgofile.txt")

}

func readFile(filname string) {
	databyte, err := os.ReadFile(filname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
