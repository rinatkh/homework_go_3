package loops

import "fmt"

// Example собирает демонстрацию по циклам в одну строку.
func Example() string {
	index, found := FindFirstNegative([]int{10, 5, -2, 7})
	maxValue, ok := MaxInSlice([]int{4, 9, 1})
	return fmt.Sprintf(
		"sum=%d factorial=%d even=%d negativeIndex=%d found=%t fizz=%v max=%d ok=%t",
		SumTo(5),
		Factorial(5),
		CountEven([]int{1, 2, 3, 4, 6}),
		index,
		found,
		FizzBuzz(5),
		maxValue,
		ok,
	)
}
