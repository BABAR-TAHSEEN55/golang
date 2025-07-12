package main

//NOTE : Channels can be buffered or Unbuffered
import "fmt"

// Cookies example some can hold 5 or 50 .Initially channels can hold only 1
func main() {
	UserCh := make(chan string)
	UserCh <- "LGG"
	user := <-UserCh
	fmt.Println(user)

}
