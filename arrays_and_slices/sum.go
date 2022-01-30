package arraysandslices

// Sum adds all numbers in a slice and returns the sum
func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// SumAll takes a slice of number slices, returns a slice with summed totals from argument slices
func SumAll(numList ...[]int) []int {
	var sum []int

	for _, valueNumList := range numList {
		sum = append(sum, Sum(valueNumList))
	}

	return sum
}

// SumAllTails takes a slice of number slices, returns a slice with summed totals from argument slices (except head values)
func SumAllTails(numList ...[]int) []int {
	var sum []int

	for _, valueNumList := range numList {
		if len(valueNumList) < 1 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(valueNumList[1:]))
		}
	}

	return sum
}
