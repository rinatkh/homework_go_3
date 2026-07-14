package arrays

import "testing"

func TestFirstOfThree(t *testing.T) {
	cases := []struct {
		name string
		in   [3]int
		want int
	}{
		{"positive", [3]int{1, 2, 3}, 1}, {"zero first", [3]int{0, 2, 3}, 0}, {"negative", [3]int{-5, 2, 3}, -5},
		{"all zero", [3]int{}, 0}, {"same", [3]int{7, 7, 7}, 7}, {"large", [3]int{1000, 2, 3}, 1000},
		{"mixed", [3]int{-1, 0, 1}, -1}, {"min sample", [3]int{-999, 4, 5}, -999}, {"last larger", [3]int{2, 3, 100}, 2},
		{"middle larger", [3]int{4, 99, 5}, 4},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FirstOfThree(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestLastOfFour(t *testing.T) {
	cases := []struct {
		name string
		in   [4]string
		want string
	}{
		{"letters", [4]string{"a", "b", "c", "d"}, "d"}, {"empty last", [4]string{"a", "b", "c", ""}, ""},
		{"all empty", [4]string{}, ""}, {"unicode", [4]string{"a", "b", "c", "Привет"}, "Привет"},
		{"emoji", [4]string{"a", "b", "c", "🙂"}, "🙂"}, {"same", [4]string{"x", "x", "x", "x"}, "x"},
		{"spaces", [4]string{"a", "b", "c", " "}, " "}, {"numbers text", [4]string{"1", "2", "3", "4"}, "4"},
		{"first only", [4]string{"go", "", "", ""}, ""}, {"long", [4]string{"a", "b", "c", "golang"}, "golang"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LastOfFour(tc.in); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestSetMiddle(t *testing.T) {
	cases := []struct {
		name  string
		in    [3]int
		value int
		want  [3]int
	}{
		{"basic", [3]int{1, 2, 3}, 9, [3]int{1, 9, 3}}, {"zero", [3]int{1, 2, 3}, 0, [3]int{1, 0, 3}},
		{"negative", [3]int{1, 2, 3}, -5, [3]int{1, -5, 3}}, {"zeros", [3]int{}, 7, [3]int{0, 7, 0}},
		{"same value", [3]int{1, 2, 3}, 2, [3]int{1, 2, 3}}, {"large", [3]int{10, 20, 30}, 1000, [3]int{10, 1000, 30}},
		{"negative input", [3]int{-1, -2, -3}, 4, [3]int{-1, 4, -3}}, {"edges zero", [3]int{0, 5, 0}, 8, [3]int{0, 8, 0}},
		{"all same", [3]int{6, 6, 6}, 1, [3]int{6, 1, 6}}, {"mixed", [3]int{-10, 0, 10}, 11, [3]int{-10, 11, 10}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := SetMiddle(input, tc.value); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if input != tc.in {
				t.Fatalf("input changed: %v", input)
			}
		})
	}
}

func TestSwapEdges(t *testing.T) {
	cases := []struct {
		name     string
		in, want [4]int
	}{
		{"basic", [4]int{1, 2, 3, 4}, [4]int{4, 2, 3, 1}}, {"zeros", [4]int{}, [4]int{}},
		{"same edges", [4]int{5, 2, 3, 5}, [4]int{5, 2, 3, 5}}, {"negative", [4]int{-1, -2, -3, -4}, [4]int{-4, -2, -3, -1}},
		{"mixed", [4]int{-10, 20, 30, 40}, [4]int{40, 20, 30, -10}}, {"large", [4]int{1000, 2, 3, 2000}, [4]int{2000, 2, 3, 1000}},
		{"middle same", [4]int{1, 7, 7, 9}, [4]int{9, 7, 7, 1}}, {"edge zero", [4]int{0, 1, 2, 3}, [4]int{3, 1, 2, 0}},
		{"last zero", [4]int{3, 1, 2, 0}, [4]int{0, 1, 2, 3}}, {"alternating", [4]int{1, -1, 1, -1}, [4]int{-1, -1, 1, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := SwapEdges(input); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if input != tc.in {
				t.Fatalf("input changed")
			}
		})
	}
}

func TestSumThree(t *testing.T) {
	cases := []struct {
		name string
		in   [3]int
		want int
	}{
		{"one two three", [3]int{1, 2, 3}, 6}, {"zeros", [3]int{}, 0}, {"negative", [3]int{-1, -2, -3}, -6},
		{"mixed zero", [3]int{-5, 0, 5}, 0}, {"same", [3]int{7, 7, 7}, 21}, {"large", [3]int{100, 200, 300}, 600},
		{"cancel", [3]int{10, -4, -6}, 0}, {"one", [3]int{1, 0, 0}, 1}, {"negative and positive", [3]int{-10, 2, 3}, -5},
		{"order irrelevant", [3]int{9, 1, 5}, 15},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumThree(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestAverageFour(t *testing.T) {
	cases := []struct {
		name string
		in   [4]int
		want int
	}{
		{"exact", [4]int{2, 4, 6, 8}, 5}, {"zeros", [4]int{}, 0}, {"truncate", [4]int{1, 2, 3, 5}, 2},
		{"negative exact", [4]int{-2, -4, -6, -8}, -5}, {"mixed", [4]int{-10, 0, 10, 20}, 5}, {"same", [4]int{7, 7, 7, 7}, 7},
		{"small", [4]int{0, 0, 0, 3}, 0}, {"negative truncate", [4]int{-1, -2, -3, -5}, -2},
		{"large", [4]int{100, 200, 300, 400}, 250}, {"cancel", [4]int{-5, -5, 5, 5}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AverageFour(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestReverseFive(t *testing.T) {
	cases := []struct {
		name     string
		in, want [5]int
	}{
		{"basic", [5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}}, {"zeros", [5]int{}, [5]int{}},
		{"one nonzero", [5]int{1, 0, 0, 0, 0}, [5]int{0, 0, 0, 0, 1}}, {"negative", [5]int{-1, -2, -3, -4, -5}, [5]int{-5, -4, -3, -2, -1}},
		{"palindrome", [5]int{1, 2, 3, 2, 1}, [5]int{1, 2, 3, 2, 1}}, {"same", [5]int{7, 7, 7, 7, 7}, [5]int{7, 7, 7, 7, 7}},
		{"mixed", [5]int{-2, 0, 3, 10, -9}, [5]int{-9, 10, 3, 0, -2}}, {"large", [5]int{10, 20, 30, 40, 50}, [5]int{50, 40, 30, 20, 10}},
		{"alternating", [5]int{1, -1, 1, -1, 1}, [5]int{1, -1, 1, -1, 1}}, {"middle only", [5]int{0, 0, 5, 0, 0}, [5]int{0, 0, 5, 0, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := ReverseFive(input); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if input != tc.in {
				t.Fatalf("input changed")
			}
		})
	}
}

func TestEqualPairs(t *testing.T) {
	cases := []struct {
		name string
		a, b [2]string
		want bool
	}{
		{"equal", [2]string{"go", "sql"}, [2]string{"go", "sql"}, true}, {"different second", [2]string{"go", "sql"}, [2]string{"go", "redis"}, false},
		{"different order", [2]string{"go", "sql"}, [2]string{"sql", "go"}, false}, {"empty", [2]string{}, [2]string{}, true},
		{"one empty", [2]string{"go", ""}, [2]string{"go", ""}, true}, {"case", [2]string{"Go", "sql"}, [2]string{"go", "sql"}, false},
		{"unicode", [2]string{"Привет", "мир"}, [2]string{"Привет", "мир"}, true}, {"spaces", [2]string{" ", "x"}, [2]string{"", "x"}, false},
		{"emoji", [2]string{"🙂", "go"}, [2]string{"🙂", "go"}, true}, {"both same single", [2]string{"x", "x"}, [2]string{"x", "x"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := EqualPairs(tc.a, tc.b); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestCopyAndSet(t *testing.T) {
	cases := []struct {
		name         string
		in           [3]int
		index, value int
		want         [3]int
	}{
		{"first", [3]int{1, 2, 3}, 0, 9, [3]int{9, 2, 3}}, {"middle", [3]int{1, 2, 3}, 1, 9, [3]int{1, 9, 3}},
		{"last", [3]int{1, 2, 3}, 2, 9, [3]int{1, 2, 9}}, {"zero value", [3]int{5, 6, 7}, 1, 0, [3]int{5, 0, 7}},
		{"negative", [3]int{5, 6, 7}, 2, -1, [3]int{5, 6, -1}}, {"same", [3]int{5, 6, 7}, 0, 5, [3]int{5, 6, 7}},
		{"zeros", [3]int{}, 2, 4, [3]int{0, 0, 4}}, {"large", [3]int{1, 1, 1}, 1, 1000, [3]int{1, 1000, 1}},
		{"negative input", [3]int{-1, -2, -3}, 0, 8, [3]int{8, -2, -3}}, {"mixed", [3]int{-5, 0, 5}, 2, 10, [3]int{-5, 0, 10}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original, changed := CopyAndSet(tc.in, tc.index, tc.value)
			if original != tc.in {
				t.Fatalf("original=%v, want %v", original, tc.in)
			}
			if changed != tc.want {
				t.Fatalf("changed=%v, want %v", changed, tc.want)
			}
		})
	}
}

func TestCountTrue(t *testing.T) {
	cases := []struct {
		name string
		in   [5]bool
		want int
	}{
		{"none", [5]bool{}, 0}, {"all", [5]bool{true, true, true, true, true}, 5}, {"first", [5]bool{true}, 1},
		{"last", [5]bool{false, false, false, false, true}, 1}, {"three", [5]bool{true, false, true, false, true}, 3},
		{"two", [5]bool{false, true, true, false, false}, 2}, {"four", [5]bool{true, true, false, true, true}, 4},
		{"middle", [5]bool{false, false, true, false, false}, 1}, {"alternating false first", [5]bool{false, true, false, true, false}, 2},
		{"alternating true first", [5]bool{true, false, true, false, true}, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountTrue(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestContainsFour(t *testing.T) {
	cases := []struct {
		name   string
		in     [4]int
		target int
		want   bool
	}{
		{"first", [4]int{1, 2, 3, 4}, 1, true}, {"middle", [4]int{1, 2, 3, 4}, 3, true}, {"last", [4]int{1, 2, 3, 4}, 4, true},
		{"missing", [4]int{1, 2, 3, 4}, 5, false}, {"zero default", [4]int{}, 0, true}, {"zero missing", [4]int{1, 2, 3, 4}, 0, false},
		{"negative", [4]int{-3, -2, -1, 0}, -2, true}, {"duplicate", [4]int{7, 7, 7, 7}, 7, true},
		{"large missing", [4]int{100, 200, 300, 400}, 500, false}, {"mixed", [4]int{-10, 0, 10, 20}, 20, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ContainsFour(tc.in, tc.target); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestMaxFour(t *testing.T) {
	cases := []struct {
		name string
		in   [4]int
		want int
	}{
		{"ascending", [4]int{1, 2, 3, 4}, 4}, {"descending", [4]int{4, 3, 2, 1}, 4}, {"all negative", [4]int{-7, -2, -9, -4}, -2},
		{"zeros", [4]int{}, 0}, {"same", [4]int{5, 5, 5, 5}, 5}, {"first max", [4]int{10, 2, 3, 4}, 10},
		{"middle max", [4]int{1, 20, 3, 4}, 20}, {"third max", [4]int{1, 2, 30, 4}, 30}, {"mixed", [4]int{-10, 0, 10, -20}, 10},
		{"large", [4]int{1000, 999, 5000, 20}, 5000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxFour(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestMainDiagonalSum(t *testing.T) {
	cases := []struct {
		name string
		in   [3][3]int
		want int
	}{
		{"one to nine", [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 15}, {"zeros", [3][3]int{}, 0},
		{"diagonal only", [3][3]int{{1, 0, 0}, {0, 2, 0}, {0, 0, 3}}, 6}, {"negative", [3][3]int{{-1, 2, 3}, {4, -5, 6}, {7, 8, -9}}, -15},
		{"off diagonal ignored", [3][3]int{{0, 9, 9}, {9, 0, 9}, {9, 9, 0}}, 0}, {"same", [3][3]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}, 6},
		{"mixed", [3][3]int{{10, 0, 0}, {0, -3, 0}, {0, 0, 5}}, 12}, {"large", [3][3]int{{100, 1, 1}, {1, 200, 1}, {1, 1, 300}}, 600},
		{"first only", [3][3]int{{7}}, 7}, {"last only", [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 8}}, 8},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MainDiagonalSum(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCornersSum(t *testing.T) {
	cases := []struct {
		name string
		in   [3][3]int
		want int
	}{
		{"one to nine", [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 20}, {"zeros", [3][3]int{}, 0},
		{"corners", [3][3]int{{1, 0, 2}, {0, 9, 0}, {3, 0, 4}}, 10}, {"center ignored", [3][3]int{{0, 0, 0}, {0, 100, 0}, {0, 0, 0}}, 0},
		{"negative", [3][3]int{{-1, 0, -2}, {0, 0, 0}, {-3, 0, -4}}, -10}, {"same", [3][3]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}, 8},
		{"mixed", [3][3]int{{10, 1, -5}, {2, 3, 4}, {7, 8, -2}}, 10}, {"first corner", [3][3]int{{9}}, 9},
		{"last corner", [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 6}}, 6}, {"large", [3][3]int{{100, 0, 200}, {0, 0, 0}, {300, 0, 400}}, 1000},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CornersSum(tc.in); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestRotateRight(t *testing.T) {
	cases := []struct {
		name     string
		in, want [4]int
	}{
		{"basic", [4]int{1, 2, 3, 4}, [4]int{4, 1, 2, 3}}, {"zeros", [4]int{}, [4]int{}},
		{"same", [4]int{7, 7, 7, 7}, [4]int{7, 7, 7, 7}}, {"negative", [4]int{-1, -2, -3, -4}, [4]int{-4, -1, -2, -3}},
		{"mixed", [4]int{-10, 0, 10, 20}, [4]int{20, -10, 0, 10}}, {"one first", [4]int{1, 0, 0, 0}, [4]int{0, 1, 0, 0}},
		{"one last", [4]int{0, 0, 0, 1}, [4]int{1, 0, 0, 0}}, {"large", [4]int{100, 200, 300, 400}, [4]int{400, 100, 200, 300}},
		{"alternating", [4]int{1, -1, 1, -1}, [4]int{-1, 1, -1, 1}}, {"ordered negative positive", [4]int{-2, -1, 0, 1}, [4]int{1, -2, -1, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := RotateRight(input); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if input != tc.in {
				t.Fatal("input changed")
			}
		})
	}
}

func TestExampleOutput(t *testing.T) {
	want := "first=10 last=d middle=[1 8 3] swap=[4 2 3 1] sum=9 reverse=[5 4 3 2 1] equal=true original=[4 5 6] changed=[4 99 6] max=-2 diagonal=15 rotate=[4 1 2 3]"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
