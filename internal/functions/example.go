package functions

import "fmt"

// Example собирает демонстрацию по функциям в одну строку.
func Example() string {
	div, ok := SafeDivide(17, 5)
	minValue, maxValue := MinMax(10, -4)
	avg, avgOK := Average(2, 4, 6)
	return fmt.Sprintf(
		"name=%s div=%d ok=%t min=%d max=%d sum=%d avg=%.1f avgOK=%t",
		FullName("Мария", "Иванова"),
		div,
		ok,
		minValue,
		maxValue,
		SumVariadic(1, 2, 3, 4),
		avg,
		avgOK,
	)
}
