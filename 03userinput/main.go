package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)

}
