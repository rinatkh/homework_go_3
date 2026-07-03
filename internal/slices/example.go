package slices

import "fmt"

// Example собирает демонстрацию по слайсам в одну строку.
func Example() string {
	numbers := []int{5, 2, 8, 2, 1}
	avg, ok := Average(numbers)
	copyOfNumbers := CopySlice(numbers)
	ReverseInPlace(copyOfNumbers)
	return fmt.Sprintf(
		"sum=%d avg=%.1f ok=%t even=%v doubled=%v unique=%v reversed=%v sorted=%v",
		Sum(numbers),
		avg,
		ok,
		FilterEven(numbers),
		MapDouble(numbers),
		Unique(numbers),
		copyOfNumbers,
		MergeAndSort([]int{3, 1}, []int{2, 4}),
	)
}
