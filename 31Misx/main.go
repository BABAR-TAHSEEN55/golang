package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type Person struct {
	Name string
	Age  uint
}

func (P *Person) CelebrateBirthday() {
	P.Age++
}

// NOTE: : 1-> Returns 1 match , 3-> Returns 3 matches , -1-> Returns all matches
// TODO : Create a Function that searches up for something
func main() {
	Ptr := &Person{"Alice", 43}
	fmt.Printf("Before Bday : %d and Name : %s", Ptr.Age, Ptr.Name)
	Ptr.CelebrateBirthday()
	fmt.Printf("After Bday : %d and Name : %s", Ptr.Age, Ptr.Name)
	para := "Mohammed babar Tahseen email:farhan.tahseen1@gmail.com but this is not correct so you'll use email:babarsmoke24@gmail.com"
	re := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-z0-9.-]+\.[a-zA-Z0-9]{2,}`)
	emails := re.FindAllString(para, -1)
	fmt.Println("Emails : ", emails)
	// ExampleInterface()
	NewExample()

}

func NewExample() {
	jsonStr := `{
			"username": "farhan",
			"exp": 1721000000,
			"isAdmin": true
		}`

	var payload map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &payload) // Convert the string into bytes as marshal expects bytes and then stores it in payload
	if err != nil {
		panic(err)
	}
	username, ok := payload["username"].(string)
	if ok {
		fmt.Println("Username : ", username)
	} else {
		fmt.Println("Username not found or doesnt exists")
	}
	admin, ok := payload["isAdmin"].(bool)
	if ok {
		fmt.Println("Admin : ", admin)
	} else {
		fmt.Println("Admin  not found or doesnt exists")
	}
	Question := map[string]uint{
		"Q1": 4,
		"Q2": 8,
	}
	for key, val := range Question {
		fmt.Printf("Key : %s and Value : %d", key, val)
	}

}
func ExampleInterface() {

	var x interface{}
	//Here Farhan will be printed as it is overrided
	x = 4
	x = "Farhan"
	fmt.Println("Value of X : ", x)
}
