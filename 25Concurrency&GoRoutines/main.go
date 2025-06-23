package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	//
	// greeter("World")
	websiteList := []string{
		"https://github.com",
		"https://google.com",
		"https://discord.com",
		"https://twitch.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // 1-> Total number o calles here it is one
	}
	wg.Wait()

}

// func greeter(s string) {
//
//	for i := 0; i < 6; i++ {
//		time.Sleep(5 * time.Millisecond)
//		fmt.Println(s)
//	}
//
// }
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in Endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
