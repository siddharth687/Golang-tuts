package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "Rb")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
