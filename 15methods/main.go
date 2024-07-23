package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	Sid := User{"Sid", "jainsiddharth687@gmail.com", true, 22}
	fmt.Println(Sid)
	fmt.Printf("Sid details are: %+v\n", Sid)
	fmt.Printf("Name is %v and email is %v\n", Sid.Name, Sid.Email)
	Sid.GetStatus()
	Sid.NewMail()
	fmt.Printf("Name is %v and email is %v\n", Sid.Name, Sid.Email)

}

// There is no concept of Inheritance in Structs
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is", u.Email)
}
