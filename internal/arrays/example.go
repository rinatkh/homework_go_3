package arrays

import "fmt"

// Example собирает демонстрацию по массивам в одну строку.
func Example() string {
	numbers := [5]int{5, 1, 9, 1, 3}
	original, changed := CopyAndSet(numbers, 0, 100)
	return fmt.Sprintf(
		"sum=%d avg=%.1f max=%d reverse=%v contains=%t original=%v changed=%v",
		Sum(numbers),
		Average(numbers),
		Max(numbers),
		Reverse(numbers),
		Contains(numbers, 9),
		original,
		changed,
	)
}
