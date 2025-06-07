package main

import "fmt"

const AdminUserName string = "Admin" //This is a constant variable , and you can say it is a global variable which means it can bev accessed any where ( Global means having first Letter capital)
// jsonweb:=30000 // This doesn't work as it is outside of a function
// var jsonweb int = 30000 //This works
func main(){
	// fmt.Println("Variables in Go")
	var username string = "JohnDoe"
	fmt.Println("Username : ",username)
	fmt.Printf("Type of Username  : %T\n",username)

	var isLoggedIn bool = true 
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of IsLoggedIn  : %T\n",isLoggedIn)

	var value int = 255 
	fmt.Println(value)
	fmt.Printf("Type of Value  : %T\n",value)
	var FloatValue float32 = 255.22
	fmt.Println(FloatValue)
	fmt.Printf("Type of FloatValue  : %T\n",FloatValue)

	var emptyValue int 
	fmt.Println(emptyValue)
	fmt.Printf("Type of emptyValue  : %T\n",emptyValue)
	var emptyusername string 
	fmt.Println(emptyusername)
	fmt.Printf("Type of emptyUsername  : %T\n",emptyusername)
	// Implicit type declaration
	var website = "www.github.com"
	fmt.Println(website)
	// website = 2 // This will cause an error because website is already treated as string
	
	//No Var style
	// This works only inside methods
	numberofUsers := 3000
	fmt.Println(numberofUsers)
	fmt.Println("Admin User Name  : ",AdminUserName)

}