package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"os"
)

func main() {
	fmt.Println("Welcome to handling files in Golang")

	fmt.Println("Enter the file name you want to create..?")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	CheckErr(err)
	CreateFile(input)

}
func CreateFile(filename string) {
	fmt.Println("Enter the Content.")
	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	file, err := os.Create(filename)
	CheckErr(err)
	//NOTE::Returns length as well as writes the string to the file
	length, err := io.WriteString(file, content)
	CheckErr(err)
	fmt.Println("length of the file : ", length)
	defer file.Close()
	ReadFiles(filename)

}
func ReadFiles(filename string) {
	databyte, err := os.ReadFile(filename)
	CheckErr(err)
	fmt.Println("Data inside the file : \n", string(databyte))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}

}
