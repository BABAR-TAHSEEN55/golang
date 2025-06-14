package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Handling Web-Requests in Golang")

	res, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Response : \n", res.Status)
	databyte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))

}
