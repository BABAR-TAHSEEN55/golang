package main

func Sum(number []int) int {
	sum := 0
	for _, number := range number {
		sum += number
	}
	return sum
}
