package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sumAll []int) {
	numberQuantities := len(numbersToSum)
	sumAll = make([]int, numberQuantities)

	for i, numbers := range numbersToSum {
		sumAll[i] = Sum(numbers)
	}

	return
}
