package main

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numberList ...[]int) []int {
	var results []int

	for _, numbers := range numberList {
		sum := Sum(numbers)

		results = append(results, sum)
	}

	return results
}

func SumTails(numberList ...[]int) []int {
	var results []int

	for _, numbers := range numberList {
		if len(numbers) == 0 {
			results = append(results, 0)
			continue
		}

		results = append(results, Sum(numbers[1:]))
	}

	return results
}
