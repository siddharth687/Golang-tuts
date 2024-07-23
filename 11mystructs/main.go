package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	Sid := User{"Sid", "jainsiddharth687@gmail.com", true, 22}
	fmt.Println(Sid)
	fmt.Printf("Sid details are: %+v\n", Sid)
	fmt.Printf("Name is %v and email is %v", Sid.Name, Sid.Email)

}

// There is no concept of Inheritance in Structs
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
