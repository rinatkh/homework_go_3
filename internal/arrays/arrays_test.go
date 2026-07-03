package arrays

import (
	"reflect"
	"testing"
)

func TestZeroArray(t *testing.T) {
	if got := ZeroArray(); got != [3]int{} {
		t.Fatalf("ZeroArray=%v", got)
	}
}

func TestWeekdays(t *testing.T) {
	got := Weekdays()
	if len(got) != 7 || got[0] != "Mon" || got[6] != "Sun" {
		t.Fatalf("Weekdays=%v", got)
	}
}

func TestSetAt(t *testing.T) {
	got := SetAt([5]int{1, 2, 3, 4, 5}, 2, 99)
	want := [5]int{1, 2, 99, 4, 5}
	if got != want {
		t.Fatalf("SetAt=%v, want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	if got := Sum([5]int{1, 2, 3, 4, 5}); got != 15 {
		t.Fatalf("Sum=%d", got)
	}
}

func TestAverage(t *testing.T) {
	if got := Average([5]int{1, 2, 3, 4, 5}); got != 3 {
		t.Fatalf("Average=%v", got)
	}
}

func TestMax(t *testing.T) {
	if got := Max([5]int{-5, -2, -9, -1, -3}); got != -1 {
		t.Fatalf("Max=%d", got)
	}
}

func TestReverse(t *testing.T) {
	want := [5]int{5, 4, 3, 2, 1}
	if got := Reverse([5]int{1, 2, 3, 4, 5}); got != want {
		t.Fatalf("Reverse=%v, want %v", got, want)
	}
}

func TestContains(t *testing.T) {
	if !Contains([5]int{1, 2, 3, 4, 5}, 4) {
		t.Fatalf("Contains returned false")
	}
}

func TestCountValue(t *testing.T) {
	if got := CountValue([5]int{1, 2, 1, 3, 1}, 1); got != 3 {
		t.Fatalf("CountValue=%d", got)
	}
}

func TestEqual(t *testing.T) {
	if !Equal([5]int{1, 2, 3, 4, 5}, [5]int{1, 2, 3, 4, 5}) {
		t.Fatalf("Equal returned false")
	}
}

func TestCopyAndSet(t *testing.T) {
	original, changed := CopyAndSet([5]int{1, 2, 3, 4, 5}, 0, 100)
	if original != [5]int{1, 2, 3, 4, 5} || changed != [5]int{100, 2, 3, 4, 5} {
		t.Fatalf("CopyAndSet original=%v changed=%v", original, changed)
	}
}

func TestFirstLast(t *testing.T) {
	first, last := FirstLast([5]int{9, 2, 3, 4, 7})
	if first != 9 || last != 7 {
		t.Fatalf("FirstLast=%d,%d", first, last)
	}
}

func TestToSlice(t *testing.T) {
	got := ToSlice([5]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ToSlice=%v, want %v", got, want)
	}
}

func TestMatrixDiagonalSum(t *testing.T) {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if got := MatrixDiagonalSum(matrix); got != 15 {
		t.Fatalf("MatrixDiagonalSum=%d", got)
	}
}

func TestCompareBySum(t *testing.T) {
	if got := CompareBySum([5]int{1, 1, 1, 1, 1}, [5]int{2, 2, 2, 2, 2}); got != "right" {
		t.Fatalf("CompareBySum=%q", got)
	}
}
