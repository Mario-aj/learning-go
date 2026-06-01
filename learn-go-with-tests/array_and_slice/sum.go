package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sumAll []int

	for _, numbers := range numbersToSum {
		sumAll = append(sumAll, Sum(numbers))
	}

	return sumAll
}
