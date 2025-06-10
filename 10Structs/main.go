package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status string
}

func main() {
	fmt.Println("Welcome to Structs in GoLang")
	// No Inheritance in Golang : No super or No Parent

	MyUser := User{"Suho Kim", "suhoKim.go.dev", 20, "Active"}
	fmt.Printf("User details : %v", MyUser)
	// For key , value use +V
	fmt.Printf("User details : %+v", MyUser)
	fmt.Printf("Type of User  : %T", MyUser)

}
