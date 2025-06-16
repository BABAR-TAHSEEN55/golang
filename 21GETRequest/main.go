package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Request in Golang")
	// HandleGetRequest("http://localhost:3000/")
	HanldePostRequest("http://localhost:3000/")

}
func HandleGetRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("Status Code : ", res.StatusCode)
	fmt.Println("Length of the Content : ", res.ContentLength)
	var ResString strings.Builder
	content, _ := io.ReadAll(res.Body)
	byteCount, _ := ResString.Write(content)
	fmt.Println("byteCount : ", byteCount)
	//NOTE : 2nd way to handle using Builder
	fmt.Println(ResString.String())
	fmt.Println("Result : ", string(content))

}
func HanldePostRequest(url string) {
	Payload := strings.NewReader(`
		{
		"name": "Suho",
		"age": 20
		}
		`)
	res, err := http.Post(url, "application/json", Payload)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
