package loops

import "fmt"

func Example() string {
	numbers := []int{1, -2, 0, 4}
	doubled := append([]int(nil), numbers...)
	DoubleInPlace(doubled)
	replaced := append([]int(nil), numbers...)
	ReplaceNegativeInPlace(replaced, 9)
	negative, found := FirstNegative(numbers)
	users := []User{{Name: "Maria", Active: true}, {Name: "Alex"}, {Name: "Ira", Active: true}}

	return fmt.Sprintf(
		"sum=%d between=%d countdown=%v factorial=%d even=%d negative=%d/%t noZero=%d limited=%d doubled=%v replaced=%v active=%d names=%v runes=%d indexes=%v repeated=%v",
		SumTo(5),
		SumBetween(3, 6),
		CountDown(4),
		Factorial(5),
		CountEven(numbers),
		negative,
		found,
		SumWithoutZeros(numbers),
		SumUntilLimit([]int{3, 4, 5}, 7),
		doubled,
		replaced,
		CountActive(users),
		ActiveNames(users),
		RuneCount("Go🙂"),
		RuneByteIndexes("Яa🙂"),
		RepeatEach([]int{1, 2}, 2),
	)
}
