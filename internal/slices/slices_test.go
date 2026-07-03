package slices

import (
	"reflect"
	"testing"
)

func TestNewSlice(t *testing.T) {
	got := NewSlice(1, 2, 3)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("NewSlice=%v", got)
	}
}

func TestAppendValue(t *testing.T) {
	got := AppendValue([]int{1, 2}, 3)
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("AppendValue=%v", got)
	}
}

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Fatalf("Sum=%d", got)
	}
}

func TestAverage(t *testing.T) {
	avg, ok := Average([]int{2, 4, 6})
	if avg != 4 || !ok {
		t.Fatalf("Average=%v,%t", avg, ok)
	}
	_, ok = Average(nil)
	if ok {
		t.Fatalf("Average nil ok=true")
	}
}

func TestFilterEven(t *testing.T) {
	want := []int{2, 4, 6}
	if got := FilterEven([]int{1, 2, 3, 4, 6}); !reflect.DeepEqual(got, want) {
		t.Fatalf("FilterEven=%v", got)
	}
}

func TestMapDouble(t *testing.T) {
	want := []int{2, 4, 6}
	if got := MapDouble([]int{1, 2, 3}); !reflect.DeepEqual(got, want) {
		t.Fatalf("MapDouble=%v", got)
	}
}

func TestFindIndex(t *testing.T) {
	if got := FindIndex([]int{4, 5, 6}, 5); got != 1 {
		t.Fatalf("FindIndex=%d", got)
	}
	if got := FindIndex([]int{4, 5, 6}, 9); got != -1 {
		t.Fatalf("FindIndex missing=%d", got)
	}
}

func TestRemoveAt(t *testing.T) {
	want := []int{1, 3}
	if got := RemoveAt([]int{1, 2, 3}, 1); !reflect.DeepEqual(got, want) {
		t.Fatalf("RemoveAt=%v", got)
	}
}

func TestInsertAt(t *testing.T) {
	want := []int{1, 9, 2, 3}
	if got := InsertAt([]int{1, 2, 3}, 1, 9); !reflect.DeepEqual(got, want) {
		t.Fatalf("InsertAt=%v", got)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3}
	copyOfOriginal := CopySlice(original)
	copyOfOriginal[0] = 100
	if original[0] == 100 {
		t.Fatalf("CopySlice must return independent copy")
	}
}

func TestReverseInPlace(t *testing.T) {
	numbers := []int{1, 2, 3}
	ReverseInPlace(numbers)
	want := []int{3, 2, 1}
	if !reflect.DeepEqual(numbers, want) {
		t.Fatalf("ReverseInPlace=%v", numbers)
	}
}

func TestUnique(t *testing.T) {
	want := []int{1, 2, 3}
	if got := Unique([]int{1, 2, 1, 3, 2}); !reflect.DeepEqual(got, want) {
		t.Fatalf("Unique=%v", got)
	}
}

func TestWindow(t *testing.T) {
	want := []int{2, 3}
	if got := Window([]int{1, 2, 3, 4}, 1, 3); !reflect.DeepEqual(got, want) {
		t.Fatalf("Window=%v", got)
	}
}

func TestChunk(t *testing.T) {
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if got := Chunk([]int{1, 2, 3, 4, 5}, 2); !reflect.DeepEqual(got, want) {
		t.Fatalf("Chunk=%v", got)
	}
}

func TestMergeAndSort(t *testing.T) {
	want := []int{1, 2, 3, 4}
	if got := MergeAndSort([]int{3, 1}, []int{4, 2}); !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeAndSort=%v", got)
	}
}
