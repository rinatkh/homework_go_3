package arrays

import "fmt"

func Example() string {
	original, changed := CopyAndSet([3]int{4, 5, 6}, 1, 99)
	return fmt.Sprintf(
		"first=%d last=%s middle=%v swap=%v sum=%d reverse=%v equal=%t original=%v changed=%v max=%d diagonal=%d rotate=%v",
		FirstOfThree([3]int{10, 20, 30}),
		LastOfFour([4]string{"a", "b", "c", "d"}),
		SetMiddle([3]int{1, 2, 3}, 8),
		SwapEdges([4]int{1, 2, 3, 4}),
		SumThree([3]int{2, 3, 4}),
		ReverseFive([5]int{1, 2, 3, 4, 5}),
		EqualPairs([2]string{"go", "sql"}, [2]string{"go", "sql"}),
		original,
		changed,
		MaxFour([4]int{-7, -2, -9, -4}),
		MainDiagonalSum([3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		RotateRight([4]int{1, 2, 3, 4}),
	)
}
