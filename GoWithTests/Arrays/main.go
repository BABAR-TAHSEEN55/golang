package main

func Sum(number []int) int {
	sum := 0
	for _, number := range number {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumber := len(numbersToSum)
	sums := make([]int, lengthOfNumber)
	for i, num := range numbersToSum {
		sums[i] = Sum(num)
	}
	return sums
}

//NOTE : go test -cover

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, num := range numbersToSum {
		if len(num) == 0 {
			sums = append(sums, 0)

		} else {

			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

//TODO : More example and 1 last refactor
//TODO : Compile the failing tes 
