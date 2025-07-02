package main

import (
	"fmt"
	"net/url"
)

const myUrl = "postgres://user:pass@host.com:5432/path?k=v#f"

func main() {

	fmt.Println("Welcome to Handling Url with golang")
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Url HOst : ", result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.User)
	fmt.Println(result.User.Password())
	partofUrl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
	}
	MakeUrl := partofUrl.String()
	fmt.Println("Url : ", MakeUrl)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
}
