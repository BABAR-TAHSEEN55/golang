package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["Py"] = "Python"
	languages["ts"] = "Typescript"

	fmt.Println("List of Languages : ", languages)
	fmt.Println(languages["JS"])
	fmt.Println(languages["Py"])
	fmt.Println(languages["ts"])
	// delete(languages, "P")
	//FIX : Why does it not throw error ? and also find loop
	fmt.Println("List of Languages : ", languages)
}
