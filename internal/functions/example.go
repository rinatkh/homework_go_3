package functions

import "fmt"

func Example() string {
	quotient, divideErr := SafeDivide(20, 4)
	users := []User{{ID: 1, Name: "Maria", Active: true}, {ID: 2, Name: "Alex"}}
	found, foundOK := FindActiveUser(users, 1)
	counter := NewCounter(5)
	accumulator := NewAccumulator(10)
	multiplier := MakeMultiplier(3)
	operation, operationOK := ChooseOperation("mul")
	operationResult := 0
	if operationOK && operation != nil {
		operationResult = operation(4, 5)
	}

	return fmt.Sprintf(
		"divide=%d/%t user=%s/%t sum=%d apply=%d allowed=%d counter=%d,%d accumulator=%d,%d multiply=%d defer=%s captured=%s deferred=%s named=%d cleanup=%v operation=%d/%t",
		quotient,
		divideErr != nil,
		found.Name,
		foundOK,
		SumAll(1, 2, 3, 4),
		Apply(6, 7, func(a, b int) int { return a + b }),
		ApplyIf(8, true, func(value int) int { return value * 2 }),
		counter(),
		counter(),
		accumulator(5),
		accumulator(-2),
		multiplier(7),
		DeferOrder(),
		CaptureDeferArgument(),
		ReadDeferredVariable(),
		IncrementNamedResult(9),
		RunWithCleanup(func() string { return "action" }, func() string { return "cleanup" }),
		operationResult,
		operationOK,
	)
}
