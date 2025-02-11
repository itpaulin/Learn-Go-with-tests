package arrays

func SumArray(numbers [5]int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumSlice(numbers []int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}
	return sum
}
