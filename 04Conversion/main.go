package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please give a rating to my Nvim Setup")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Print("Thanks for the rating , ", input)

	numRatin, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error  : ", err)
	} else {
		fmt.Println("Adding +1 to your rating , ", numRatin+1)
	}

}
