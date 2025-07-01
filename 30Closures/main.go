package main

import "fmt"

func main() {
	Result := InitSeq()
	fmt.Println(Result())
	fmt.Println(Result())
	fmt.Println(Result())

}
func InitSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
