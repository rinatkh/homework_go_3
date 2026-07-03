package loops

import (
	"reflect"
	"testing"
)

func TestSumTo(t *testing.T) {
	tests := []struct{ n, want int }{{5, 15}, {1, 1}, {0, 0}, {-3, 0}}
	for _, tt := range tests {
		if got := SumTo(tt.n); got != tt.want {
			t.Fatalf("SumTo(%d)=%d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{{5, 120}, {0, 1}, {1, 1}, {-1, 0}}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Fatalf("Factorial(%d)=%d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCountEven(t *testing.T) {
	got := CountEven([]int{1, 2, 3, 4, 6})
	if got != 3 {
		t.Fatalf("CountEven=%d, want 3", got)
	}
}

func TestFindFirstNegative(t *testing.T) {
	index, found := FindFirstNegative([]int{4, 0, -2, -5})
	if index != 2 || !found {
		t.Fatalf("FindFirstNegative got (%d,%t), want (2,true)", index, found)
	}
	index, found = FindFirstNegative([]int{4, 0, 2})
	if index != -1 || found {
		t.Fatalf("FindFirstNegative got (%d,%t), want (-1,false)", index, found)
	}
}

func TestSkipMultiplesOfThree(t *testing.T) {
	want := []int{1, 2, 4, 5, 7}
	if got := SkipMultiplesOfThree(7); !reflect.DeepEqual(got, want) {
		t.Fatalf("SkipMultiplesOfThree=%v, want %v", got, want)
	}
}

func TestMultiplicationRow(t *testing.T) {
	want := []int{4, 8, 12, 16}
	if got := MultiplicationRow(4, 4); !reflect.DeepEqual(got, want) {
		t.Fatalf("MultiplicationRow=%v, want %v", got, want)
	}
}

func TestReverseStringByRunes(t *testing.T) {
	if got := ReverseStringByRunes("Go🙂"); got != "🙂oG" {
		t.Fatalf("ReverseStringByRunes=%q", got)
	}
}

func TestCountRunes(t *testing.T) {
	got := CountRunes("go g")
	if got['g'] != 2 || got['o'] != 1 || got[' '] != 1 {
		t.Fatalf("CountRunes=%v", got)
	}
}

func TestFizzBuzz(t *testing.T) {
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if got := FizzBuzz(15); !reflect.DeepEqual(got, want) {
		t.Fatalf("FizzBuzz=%v, want %v", got, want)
	}
}

func TestTriangleRows(t *testing.T) {
	want := []string{"*", "**", "***"}
	if got := TriangleRows(3); !reflect.DeepEqual(got, want) {
		t.Fatalf("TriangleRows=%v, want %v", got, want)
	}
}

func TestSumUntilLimit(t *testing.T) {
	if got := SumUntilLimit([]int{2, 3, 10, 1}, 6); got != 5 {
		t.Fatalf("SumUntilLimit=%d, want 5", got)
	}
}

func TestFlatten(t *testing.T) {
	want := []int{1, 2, 3, 4, 5}
	if got := Flatten([][]int{{1, 2}, {3}, {4, 5}}); !reflect.DeepEqual(got, want) {
		t.Fatalf("Flatten=%v, want %v", got, want)
	}
}

func TestMaxInSlice(t *testing.T) {
	got, ok := MaxInSlice([]int{-3, -1, -7})
	if got != -1 || !ok {
		t.Fatalf("MaxInSlice=%d,%t", got, ok)
	}
	_, ok = MaxInSlice(nil)
	if ok {
		t.Fatalf("MaxInSlice(nil) ok=true")
	}
}

func TestUniquePreserveOrder(t *testing.T) {
	want := []int{3, 1, 2, 4}
	if got := UniquePreserveOrder([]int{3, 1, 3, 2, 1, 4}); !reflect.DeepEqual(got, want) {
		t.Fatalf("UniquePreserveOrder=%v, want %v", got, want)
	}
}

func TestRepeatString(t *testing.T) {
	if got := RepeatString("go", 3); got != "gogogo" {
		t.Fatalf("RepeatString=%q", got)
	}
}
