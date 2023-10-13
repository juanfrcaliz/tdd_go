package arraysslices

// Sums up all the integers in a slice or array.
func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum += number
	}

	return
}

// Sums up the integers in each slice or array.
func SumAll(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		result = append(result, Sum(slice))
	}

	return
}

// Sums up the tails of each slice. A tail is all numbers except the first one.
func SumAllTails(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		if len(slice) > 0 {
			tail := slice[1:]
			result = append(result, Sum(tail))
		} else {
			result = append(result, 0)
		}
	}
	return
}
