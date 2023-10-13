package arraysslices

// Sums up all the integers in an array.
func Sum(arr []int) (sum int) {
	for _, number := range arr {
		sum += number
	}

	return
}
