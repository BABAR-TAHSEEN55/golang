package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Tarot Dice Game")

	DicedNumber := rand.Intn(6) + 1
	fmt.Println("You rolled:", DicedNumber)
	switch DicedNumber {
	case 1:
		fmt.Println("Card: The Fool – New beginnings, spontaneity, a leap of faith.")
	case 2:
		fmt.Println("Card: The Magician – Manifestation, resourcefulness, power.")
	case 3:
		fmt.Println("Card: The High Priestess – Intuition, mystery, the subconscious.")
	case 4:
		fmt.Println("Card: The Emperor – Authority, structure, stability.")
	case 5:
		fmt.Println("Card: The Lovers – Relationships, choices, harmony.")
	case 6:
		fmt.Println("Card: The Wheel of Fortune – Fate, change, cycles.")
	default:
		fmt.Println("Unknown card.")
	}
}
