package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to Structs in GoLang")
	// No Inheritance in Golang : No super or No Parent

	MyUser := User{"Suho Kim", "suhoKim.go.dev", 20, true}
	fmt.Printf("User details : %v", MyUser)
	// For key , value use +V
	fmt.Printf("User details : %+v", MyUser)
	fmt.Printf("Type of User  : %T\n", MyUser)
	MyUser.GetStatus()
	MyUser.NewEmail()
	fmt.Printf("User details : %v", MyUser)

}
func (user User) GetStatus() {
	fmt.Println("Is User Active ? ", user.Status)
}
func (user User) ss() {

	fmt.Println("Is User Active ? ", user.Status)
}
func (user *User) NewEmail() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter new Email...")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	user.Email = input
	fmt.Println("New Email : ", input)
}

//TODO : Add a regex for email ending
