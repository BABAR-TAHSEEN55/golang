package main

import (
	"fmt"
	"strings"
)

//TODO : Write Example Repeat
// NOTE : Strings are immutable meaning every concatenation such as in our Repeat is a heavy operation.In order to reduce it , String builder pkg is used

func Repeat(al string, iterate int) string {
	var repeated strings.Builder
	for i := 0; i < iterate; i++ {
		repeated.WriteString(al)

	}
	return repeated.String()

}

func main() {
	fmt.Println(Repeat("a", 4))

}
