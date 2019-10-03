package datamath

// Sum calculates the total from an array of numbers
func Sum(numbers []int) int {
	var sum int

	for _, v := range numbers {
		sum += v
	}

	return sum
}
