package arrays

func Sum(numbers []int) int {
	acc := 0

	for _, number := range numbers {
		acc += number
	}

	return acc
}
