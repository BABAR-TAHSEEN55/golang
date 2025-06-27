package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	country := make(map[string]bool)
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
	country["IND"] = true
	country["PAK"] = false
	for k, _ := range country {
		// country[k] = vis

		if !country[k] {
			country[k] = true
			fmt.Println("Unvisited Has been Visited")
			fmt.Println("Unvisted Country was : ", k)
		} else {
			fmt.Println("Already Visited")
		}
	}

}
