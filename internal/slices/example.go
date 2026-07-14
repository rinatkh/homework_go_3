package slices

import "fmt"

func Example() string {
	source := []int{10, 20, 30, 40}
	_, regularLen, regularCap := SliceInfo(source, 1, 3)
	_, limitedLen, limitedCap := FullSliceInfo(source, 1, 3, 3)

	mutated := MutateWindowInFunction([]int{10, 20, 30, 40}, 1, 3, 999)
	appendedSource, appendedPart := AppendWindowInFunction([]int{10, 20, 30, 40}, 1, 3, 777)
	limitedSource, limitedPart := AppendLimitedWindowInFunction([]int{10, 20, 30, 40}, 1, 3, 777)

	original := make([]int, 2, 5)
	original[0], original[1] = 1, 2
	independent := AppendIndependent(original, 3, 4)

	first, firstOK := First(source)
	last, lastOK := Last([]string{"go", "sql"})

	return fmt.Sprintf(
		"first=%d/%t last=%s/%t regular=%d/%d limited=%d/%d mutated=%v appendSource=%v appendPart=%v limitedSource=%v limitedPart=%v grow=%t kind=%s independent=%v original=%v",
		first,
		firstOK,
		last,
		lastOK,
		regularLen,
		regularCap,
		limitedLen,
		limitedCap,
		mutated,
		appendedSource,
		appendedPart,
		limitedSource,
		limitedPart,
		CanAppendWithoutGrow(original, 3),
		SliceKind(original),
		independent,
		original,
	)
}
