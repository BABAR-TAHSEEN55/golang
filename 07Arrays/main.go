package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")
	var FruitList [4]string
	FruitList[0] = "Apple"
	FruitList[1] = "Mango"
	FruitList[3] = "DragronFruit"
	fmt.Println("FruitList : ", FruitList)
	fmt.Println("FruitList : ", FruitList[0])
	fmt.Println("FruitList : ", len(FruitList))

	var VegList = [3]string{"Potato", "GreenLeaf", "RedLeaf"}
	fmt.Println("Veglist: ", VegList)
}
