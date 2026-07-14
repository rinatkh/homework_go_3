package loops

import (
	"reflect"
	"testing"
)

func TestSumTo(t *testing.T) {
	cases := []struct {
		name    string
		n, want int
	}{
		{"negative", -5, 0}, {"minus one", -1, 0}, {"zero", 0, 0}, {"one", 1, 1}, {"two", 2, 3},
		{"three", 3, 6}, {"five", 5, 15}, {"ten", 10, 55}, {"twenty", 20, 210}, {"hundred", 100, 5050},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumTo(tc.n); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestSumBetween(t *testing.T) {
	cases := []struct {
		name             string
		start, end, want int
	}{
		{"one to five", 1, 5, 15}, {"same positive", 3, 3, 3}, {"same zero", 0, 0, 0},
		{"start greater", 5, 1, 0}, {"negative range", -3, -1, -6}, {"cross zero", -2, 2, 0},
		{"zero to three", 0, 3, 6}, {"three to six", 3, 6, 18}, {"minus one to one", -1, 1, 0},
		{"ten to twelve", 10, 12, 33},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumBetween(tc.start, tc.end); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCountDown(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{"negative", -2, []int{}}, {"zero", 0, []int{}}, {"one", 1, []int{1}},
		{"two", 2, []int{2, 1}}, {"three", 3, []int{3, 2, 1}}, {"four", 4, []int{4, 3, 2, 1}},
		{"five", 5, []int{5, 4, 3, 2, 1}}, {"six", 6, []int{6, 5, 4, 3, 2, 1}},
		{"seven", 7, []int{7, 6, 5, 4, 3, 2, 1}}, {"ten", 10, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := CountDown(tc.n)
			if got == nil {
				t.Fatal("result must be non-nil")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		name    string
		n, want int
	}{
		{"negative", -3, 0}, {"minus one", -1, 0}, {"zero", 0, 1}, {"one", 1, 1}, {"two", 2, 2},
		{"three", 3, 6}, {"four", 4, 24}, {"five", 5, 120}, {"six", 6, 720}, {"seven", 7, 5040},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Factorial(tc.n); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestCountEven(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"one even", []int{2}, 1}, {"one odd", []int{1}, 0},
		{"zero", []int{0}, 1}, {"mixed", []int{1, 2, 3, 4}, 2}, {"negative", []int{-4, -3, -2, -1}, 2},
		{"all even", []int{2, 4, 6, 8}, 4}, {"all odd", []int{1, 3, 5, 7}, 0}, {"duplicates", []int{2, 2, 3, 3, 0}, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountEven(tc.items); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestFirstNegative(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  int
		ok    bool
	}{
		{"nil", nil, 0, false}, {"empty", []int{}, 0, false}, {"none", []int{0, 1, 2}, 0, false},
		{"first", []int{-1, 2, 3}, -1, true}, {"middle", []int{1, -2, 3}, -2, true}, {"last", []int{1, 2, -3}, -3, true},
		{"several", []int{1, -2, -3}, -2, true}, {"all negative", []int{-5, -4, -3}, -5, true},
		{"zero before", []int{0, -7}, -7, true}, {"large", []int{100, -999, 1}, -999, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := FirstNegative(tc.items)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %d/%t, want %d/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestSumWithoutZeros(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"only zeros", []int{0, 0, 0}, 0},
		{"positive", []int{1, 2, 3}, 6}, {"with zeros", []int{1, 0, 2, 0, 3}, 6}, {"negative", []int{-1, -2, -3}, -6},
		{"mixed", []int{-2, 0, 5, -1}, 2}, {"one", []int{7}, 7}, {"zero and negative", []int{0, -5, 0}, -5},
		{"cancel", []int{10, 0, -10}, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumWithoutZeros(tc.items); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestSumUntilLimit(t *testing.T) {
	cases := []struct {
		name        string
		items       []int
		limit, want int
	}{
		{"nil", nil, 10, 0}, {"empty", []int{}, 10, 0}, {"negative limit", []int{1, 2}, -1, 0},
		{"exact", []int{3, 4, 5}, 7, 7}, {"stop first", []int{8, 1}, 7, 0}, {"all fit", []int{1, 2, 3}, 10, 6},
		{"equal one", []int{5, 1}, 5, 5}, {"zero values", []int{0, 0, 1}, 1, 1},
		{"negative lowers sum", []int{5, -2, 4}, 7, 7}, {"order matters", []int{4, 4, 1}, 8, 8},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumUntilLimit(tc.items, tc.limit); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestDoubleInPlace(t *testing.T) {
	cases := []struct {
		name        string
		items, want []int
	}{
		{"nil", nil, nil}, {"empty", []int{}, []int{}}, {"one", []int{2}, []int{4}},
		{"three", []int{1, 2, 3}, []int{2, 4, 6}}, {"zero", []int{0}, []int{0}}, {"negative", []int{-1, -2}, []int{-2, -4}},
		{"mixed", []int{-2, 0, 3}, []int{-4, 0, 6}}, {"large", []int{100, 200}, []int{200, 400}},
		{"duplicates", []int{7, 7, 7}, []int{14, 14, 14}}, {"cancel irrelevant", []int{-5, 5}, []int{-10, 10}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			DoubleInPlace(tc.items)
			if !reflect.DeepEqual(tc.items, tc.want) {
				t.Fatalf("got %v, want %v", tc.items, tc.want)
			}
		})
	}
}

func TestReplaceNegativeInPlace(t *testing.T) {
	cases := []struct {
		name        string
		items       []int
		replacement int
		want        []int
	}{
		{"nil", nil, 0, nil}, {"empty", []int{}, 0, []int{}}, {"none", []int{0, 1, 2}, 9, []int{0, 1, 2}},
		{"all", []int{-1, -2, -3}, 0, []int{0, 0, 0}}, {"mixed", []int{-1, 0, 2, -3}, 7, []int{7, 0, 2, 7}},
		{"negative replacement", []int{-1, 2}, -9, []int{-9, 2}}, {"first", []int{-5, 1}, 8, []int{8, 1}},
		{"last", []int{1, -5}, 8, []int{1, 8}}, {"large", []int{-100, 200, -300}, 1, []int{1, 200, 1}},
		{"zero unchanged", []int{0, -1, 0}, 5, []int{0, 5, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ReplaceNegativeInPlace(tc.items, tc.replacement)
			if !reflect.DeepEqual(tc.items, tc.want) {
				t.Fatalf("got %v, want %v", tc.items, tc.want)
			}
		})
	}
}

func TestCountActive(t *testing.T) {
	cases := []struct {
		name  string
		users []User
		want  int
	}{
		{"nil", nil, 0}, {"empty", []User{}, 0}, {"one active", []User{{Name: "M", Active: true}}, 1},
		{"one inactive", []User{{Name: "A"}}, 0}, {"mixed", []User{{Active: true}, {}, {Active: true}}, 2},
		{"all active", []User{{Active: true}, {Active: true}, {Active: true}}, 3}, {"all inactive", []User{{}, {}, {}}, 0},
		{"names irrelevant", []User{{Name: "M", Active: true}, {Name: "A", Active: true}}, 2},
		{"five", []User{{Active: true}, {}, {Active: true}, {}, {Active: true}}, 3},
		{"unicode", []User{{Name: "Мария", Active: true}, {Name: "Ира"}}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountActive(tc.users); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestActiveNames(t *testing.T) {
	cases := []struct {
		name  string
		users []User
		want  []string
	}{
		{"nil", nil, []string{}}, {"empty", []User{}, []string{}}, {"one active", []User{{Name: "M", Active: true}}, []string{"M"}},
		{"one inactive", []User{{Name: "A"}}, []string{}}, {"mixed", []User{{Name: "M", Active: true}, {Name: "A"}, {Name: "I", Active: true}}, []string{"M", "I"}},
		{"all active", []User{{Name: "A", Active: true}, {Name: "B", Active: true}}, []string{"A", "B"}},
		{"all inactive", []User{{Name: "A"}, {Name: "B"}}, []string{}}, {"empty name active", []User{{Name: "", Active: true}}, []string{""}},
		{"unicode", []User{{Name: "Мария", Active: true}, {Name: "Ира", Active: true}}, []string{"Мария", "Ира"}},
		{"order", []User{{Name: "3", Active: true}, {Name: "1", Active: true}, {Name: "2", Active: true}}, []string{"3", "1", "2"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ActiveNames(tc.users)
			if got == nil {
				t.Fatal("result must be non-nil")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRuneCount(t *testing.T) {
	cases := []struct {
		name, text string
		want       int
	}{
		{"empty", "", 0}, {"ascii", "hello", 5}, {"cyrillic", "Привет", 6}, {"emoji", "🙂", 1},
		{"mixed", "Go🙂", 3}, {"spaces", "a b", 3}, {"newline", "a\nb", 3}, {"two emoji", "🙂🙃", 2},
		{"combination", "Яa🙂", 3}, {"digits", "12345", 5},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RuneCount(tc.text); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestRuneByteIndexes(t *testing.T) {
	cases := []struct {
		name, text string
		want       []int
	}{
		{"empty", "", []int{}}, {"ascii", "abc", []int{0, 1, 2}}, {"cyrillic", "Яб", []int{0, 2}},
		{"emoji", "🙂", []int{0}}, {"mixed", "Яa🙂", []int{0, 2, 3}}, {"space", "a b", []int{0, 1, 2}},
		{"two emoji", "🙂🙃", []int{0, 4}}, {"latin cyrillic", "aЯb", []int{0, 1, 3}},
		{"newline", "a\nЯ", []int{0, 1, 2}}, {"digits", "12", []int{0, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RuneByteIndexes(tc.text)
			if got == nil {
				t.Fatal("result must be non-nil")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestRepeatEach(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		times int
		want  []int
	}{
		{"nil", nil, 2, []int{}}, {"empty", []int{}, 2, []int{}}, {"zero times", []int{1, 2}, 0, []int{}},
		{"negative times", []int{1, 2}, -1, []int{}}, {"once", []int{1, 2}, 1, []int{1, 2}},
		{"twice", []int{1, 2}, 2, []int{1, 1, 2, 2}}, {"three", []int{3}, 3, []int{3, 3, 3}},
		{"negative values", []int{-1, 0}, 2, []int{-1, -1, 0, 0}}, {"duplicates", []int{7, 7}, 2, []int{7, 7, 7, 7}},
		{"three items", []int{1, 2, 3}, 2, []int{1, 1, 2, 2, 3, 3}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RepeatEach(tc.items, tc.times)
			if got == nil {
				t.Fatal("result must be non-nil")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestExampleOutput(t *testing.T) {
	want := "sum=15 between=18 countdown=[4 3 2 1] factorial=120 even=3 negative=-2/true noZero=3 limited=7 doubled=[2 -4 0 8] replaced=[1 9 0 4] active=2 names=[Maria Ira] runes=3 indexes=[0 2 3] repeated=[1 1 2 2]"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
