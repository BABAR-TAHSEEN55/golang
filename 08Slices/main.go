package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to Slices")
	var Anime = []string{"Goku", "Luffy", "Ichigo"}
	fmt.Printf("Type of Anime : %T\n", Anime)
	Anime = append(Anime, "Klein", "Super")
	Anime = append(Anime[1:])
	fmt.Println(Anime)

	HighScores := make([]int, 4)
	HighScores[0] = 234
	HighScores[1] = 554
	HighScores[2] = 123
	HighScores[3] = 11
	fmt.Println(HighScores)
	fmt.Println(sort.IntsAreSorted(HighScores))
	HighScores = append(HighScores, 324, 111, 555)
	sort.Ints(HighScores)
	fmt.Println(HighScores)
	fmt.Println(sort.IntsAreSorted(HighScores))
	// TODO : Find out a way to Sort in decreaing Order
	sort.Ints(HighScores)
	var courses = []string{"Js", "React", "Python", "Ruby"}
	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println("Courses : ", courses)

}

//NOTE : This ( := ) is used for Short variable declartion within  functions only
