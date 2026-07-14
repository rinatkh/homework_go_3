package slices

import (
	"reflect"
	"testing"
)

func TestFirst(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  int
		ok    bool
	}{
		{"nil", nil, 0, false}, {"empty", []int{}, 0, false}, {"one", []int{7}, 7, true},
		{"positive", []int{1, 2, 3}, 1, true}, {"zero first", []int{0, 2, 3}, 0, true}, {"negative", []int{-5, 2, 3}, -5, true},
		{"last larger", []int{1, 2, 100}, 1, true}, {"same", []int{9, 9}, 9, true}, {"large", []int{1000}, 1000, true},
		{"mixed", []int{-1, 0, 1}, -1, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := First(tc.items)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %d/%t, want %d/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestLast(t *testing.T) {
	cases := []struct {
		name  string
		items []string
		want  string
		ok    bool
	}{
		{"nil", nil, "", false}, {"empty", []string{}, "", false}, {"one", []string{"go"}, "go", true},
		{"two", []string{"go", "sql"}, "sql", true}, {"empty last", []string{"go", ""}, "", true},
		{"unicode", []string{"a", "Привет"}, "Привет", true}, {"emoji", []string{"a", "🙂"}, "🙂", true},
		{"space", []string{"a", " "}, " ", true}, {"same", []string{"x", "x"}, "x", true},
		{"long", []string{"a", "long text"}, "long text", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := Last(tc.items)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %q/%t, want %q/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestSliceInfo(t *testing.T) {
	cases := []struct {
		name             string
		items            []int
		low, high        int
		want             []int
		wantLen, wantCap int
	}{
		{"middle literal", []int{10, 20, 30, 40}, 1, 3, []int{20, 30}, 2, 3},
		{"all literal", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}, 3, 3},
		{"empty at start", []int{1, 2, 3}, 0, 0, []int{}, 0, 3},
		{"empty at end", []int{1, 2, 3}, 3, 3, []int{}, 0, 0},
		{"one", []int{1, 2, 3}, 1, 2, []int{2}, 1, 2},
		{"extra cap", withCap([]int{1, 2, 3, 4}, 8), 1, 3, []int{2, 3}, 2, 7},
		{"extra cap all", withCap([]int{1, 2}, 5), 0, 2, []int{1, 2}, 2, 5},
		{"negative values", []int{-3, -2, -1, 0}, 2, 4, []int{-1, 0}, 2, 2},
		{"zero values", make([]int, 4, 7), 1, 2, []int{0}, 1, 6},
		{"tail", withCap([]int{5, 6, 7, 8, 9}, 10), 3, 5, []int{8, 9}, 2, 7},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			part, length, capacity := SliceInfo(tc.items, tc.low, tc.high)
			if !reflect.DeepEqual(part, tc.want) || length != tc.wantLen || capacity != tc.wantCap {
				t.Fatalf("got %v %d/%d, want %v %d/%d", part, length, capacity, tc.want, tc.wantLen, tc.wantCap)
			}
		})
	}
}

func TestFullSliceInfo(t *testing.T) {
	cases := []struct {
		name             string
		items            []int
		low, high, max   int
		want             []int
		wantLen, wantCap int
	}{
		{"limit at high", []int{10, 20, 30, 40}, 1, 3, 3, []int{20, 30}, 2, 2},
		{"keep full cap", []int{10, 20, 30, 40}, 1, 3, 4, []int{20, 30}, 2, 3},
		{"all", []int{1, 2, 3}, 0, 3, 3, []int{1, 2, 3}, 3, 3},
		{"empty start", []int{1, 2, 3}, 0, 0, 0, []int{}, 0, 0},
		{"empty with cap", []int{1, 2, 3}, 0, 0, 2, []int{}, 0, 2},
		{"one limited", []int{1, 2, 3}, 1, 2, 2, []int{2}, 1, 1},
		{"extra source cap", withCap([]int{1, 2, 3, 4}, 8), 1, 3, 5, []int{2, 3}, 2, 4},
		{"tail", []int{-3, -2, -1, 0}, 2, 4, 4, []int{-1, 0}, 2, 2},
		{"middle max one extra", withCap([]int{5, 6, 7, 8}, 10), 1, 2, 3, []int{6}, 1, 2},
		{"empty end", withCap([]int{1, 2, 3}, 6), 3, 3, 5, []int{}, 0, 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			part, length, capacity := FullSliceInfo(tc.items, tc.low, tc.high, tc.max)
			if !reflect.DeepEqual(part, tc.want) || length != tc.wantLen || capacity != tc.wantCap {
				t.Fatalf("got %v %d/%d, want %v %d/%d", part, length, capacity, tc.want, tc.wantLen, tc.wantCap)
			}
		})
	}
}

func TestChangeFirst(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		value int
		want  []int
	}{
		{"nil", nil, 9, nil}, {"empty", []int{}, 9, []int{}}, {"one", []int{1}, 9, []int{9}},
		{"three", []int{1, 2, 3}, 7, []int{7, 2, 3}}, {"zero", []int{1, 2}, 0, []int{0, 2}},
		{"negative", []int{1, 2}, -5, []int{-5, 2}}, {"same", []int{4, 5}, 4, []int{4, 5}},
		{"large", []int{1, 2, 3}, 1000, []int{1000, 2, 3}}, {"subslice", []int{20, 30}, 99, []int{99, 30}},
		{"mixed", []int{-1, 0, 1}, 8, []int{8, 0, 1}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ChangeFirst(tc.items, tc.value)
			if !reflect.DeepEqual(tc.items, tc.want) {
				t.Fatalf("got %v, want %v", tc.items, tc.want)
			}
		})
	}
}

func TestMutateWindowInFunction(t *testing.T) {
	cases := []struct {
		name             string
		items            []int
		low, high, value int
		want             []int
	}{
		{"middle", []int{10, 20, 30, 40}, 1, 3, 999, []int{10, 999, 30, 40}},
		{"first", []int{1, 2, 3}, 0, 2, 7, []int{7, 2, 3}}, {"last window", []int{1, 2, 3}, 2, 3, 8, []int{1, 2, 8}},
		{"empty window", []int{1, 2, 3}, 1, 1, 9, []int{1, 2, 3}}, {"all", []int{1, 2}, 0, 2, 0, []int{0, 2}},
		{"negative", []int{-1, -2, -3}, 1, 3, 5, []int{-1, 5, -3}}, {"same", []int{4, 5, 6}, 1, 2, 5, []int{4, 5, 6}},
		{"large", []int{1, 2, 3, 4, 5}, 3, 5, 100, []int{1, 2, 3, 100, 5}},
		{"one item", []int{7}, 0, 1, 9, []int{9}}, {"zero values", []int{0, 0, 0}, 1, 3, -1, []int{0, -1, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := MutateWindowInFunction(tc.items, tc.low, tc.high, tc.value)
			if !reflect.DeepEqual(got, tc.want) || !reflect.DeepEqual(tc.items, tc.want) {
				t.Fatalf("got %v items %v, want %v", got, tc.items, tc.want)
			}
		})
	}
}

func TestAppendWindowInFunction(t *testing.T) {
	cases := []struct {
		name                 string
		items                []int
		low, high, value     int
		wantSource, wantPart []int
	}{
		{"overwrite next", []int{10, 20, 30, 40}, 1, 3, 777, []int{10, 20, 30, 777}, []int{20, 30, 777}},
		{"append at end allocates", []int{1, 2, 3}, 1, 3, 9, []int{1, 2, 3}, []int{2, 3, 9}},
		{"from start overwrite", []int{1, 2, 3, 4}, 0, 2, 8, []int{1, 2, 8, 4}, []int{1, 2, 8}},
		{"one then overwrite", []int{1, 2, 3}, 1, 2, 7, []int{1, 2, 7}, []int{2, 7}},
		{"empty window overwrite", []int{1, 2, 3}, 1, 1, 6, []int{1, 6, 3}, []int{6}},
		{"empty at end allocates", []int{1, 2, 3}, 3, 3, 6, []int{1, 2, 3}, []int{6}},
		{"negative", []int{-1, -2, -3, -4}, 1, 3, 0, []int{-1, -2, -3, 0}, []int{-2, -3, 0}},
		{"zero", []int{0, 0, 0, 0}, 2, 3, 5, []int{0, 0, 0, 5}, []int{0, 5}},
		{"full slice allocates", []int{5, 6}, 0, 2, 7, []int{5, 6}, []int{5, 6, 7}},
		{"extra source cap beyond len", withCap([]int{1, 2}, 5), 0, 2, 3, []int{1, 2}, []int{1, 2, 3}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source, part := AppendWindowInFunction(tc.items, tc.low, tc.high, tc.value)
			if !reflect.DeepEqual(source, tc.wantSource) || !reflect.DeepEqual(part, tc.wantPart) {
				t.Fatalf("source=%v part=%v, want %v %v", source, part, tc.wantSource, tc.wantPart)
			}
		})
	}
}

func TestAppendLimitedWindowInFunction(t *testing.T) {
	cases := []struct {
		name                 string
		items                []int
		low, high, value     int
		wantSource, wantPart []int
	}{
		{"middle", []int{10, 20, 30, 40}, 1, 3, 777, []int{10, 20, 30, 40}, []int{20, 30, 777}},
		{"start", []int{1, 2, 3, 4}, 0, 2, 8, []int{1, 2, 3, 4}, []int{1, 2, 8}},
		{"end", []int{1, 2, 3}, 1, 3, 9, []int{1, 2, 3}, []int{2, 3, 9}},
		{"one", []int{1, 2, 3}, 1, 2, 7, []int{1, 2, 3}, []int{2, 7}},
		{"empty middle", []int{1, 2, 3}, 1, 1, 6, []int{1, 2, 3}, []int{6}},
		{"empty end", []int{1, 2, 3}, 3, 3, 6, []int{1, 2, 3}, []int{6}},
		{"negative", []int{-1, -2, -3, -4}, 1, 3, 0, []int{-1, -2, -3, -4}, []int{-2, -3, 0}},
		{"zeros", []int{0, 0, 0, 0}, 2, 3, 5, []int{0, 0, 0, 0}, []int{0, 5}},
		{"full", []int{5, 6}, 0, 2, 7, []int{5, 6}, []int{5, 6, 7}},
		{"extra cap", withCap([]int{1, 2, 3}, 8), 0, 2, 4, []int{1, 2, 3}, []int{1, 2, 4}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			source, part := AppendLimitedWindowInFunction(tc.items, tc.low, tc.high, tc.value)
			if !reflect.DeepEqual(source, tc.wantSource) || !reflect.DeepEqual(part, tc.wantPart) {
				t.Fatalf("source=%v part=%v, want %v %v", source, part, tc.wantSource, tc.wantPart)
			}
		})
	}
}

func TestClone(t *testing.T) {
	cases := []struct {
		name  string
		items []int
	}{
		{"nil", nil}, {"empty", []int{}}, {"one", []int{1}}, {"three", []int{1, 2, 3}},
		{"negative", []int{-1, -2}}, {"zeros", []int{0, 0, 0}}, {"same", []int{7, 7, 7}},
		{"mixed", []int{-1, 0, 1}}, {"large", []int{1000, 2000}}, {"extra cap", withCap([]int{1, 2}, 5)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Clone(tc.items)
			if !reflect.DeepEqual(got, tc.items) {
				t.Fatalf("got %v, want %v", got, tc.items)
			}
			if tc.items == nil {
				if got != nil {
					t.Fatal("nil source must produce nil clone")
				}
				return
			}
			if got == nil {
				t.Fatal("non-nil source produced nil")
			}
			if len(got) > 0 {
				before := tc.items[0]
				got[0]++
				if tc.items[0] != before {
					t.Fatal("clone shares backing array")
				}
			}
		})
	}
}

func TestChangeClone(t *testing.T) {
	cases := []struct {
		name         string
		items        []int
		index, value int
		wantClone    []int
	}{
		{"first", []int{1, 2, 3}, 0, 9, []int{9, 2, 3}}, {"middle", []int{1, 2, 3}, 1, 9, []int{1, 9, 3}},
		{"last", []int{1, 2, 3}, 2, 9, []int{1, 2, 9}}, {"one", []int{7}, 0, 8, []int{8}},
		{"zero", []int{1, 2}, 1, 0, []int{1, 0}}, {"negative", []int{1, 2}, 0, -1, []int{-1, 2}},
		{"same", []int{4, 5}, 1, 5, []int{4, 5}}, {"mixed", []int{-1, 0, 1}, 1, 7, []int{-1, 7, 1}},
		{"large", []int{100, 200, 300}, 2, 999, []int{100, 200, 999}}, {"extra cap", withCap([]int{1, 2, 3}, 8), 1, 4, []int{1, 4, 3}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := append([]int(nil), tc.items...)
			source, clone := ChangeClone(tc.items, tc.index, tc.value)
			if !reflect.DeepEqual(source, before) || !reflect.DeepEqual(tc.items, before) {
				t.Fatalf("source changed: %v", source)
			}
			if !reflect.DeepEqual(clone, tc.wantClone) {
				t.Fatalf("clone=%v, want %v", clone, tc.wantClone)
			}
		})
	}
}

func TestAppendOne(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		value int
		want  []int
	}{
		{"nil", nil, 1, []int{1}}, {"empty", []int{}, 2, []int{2}}, {"one", []int{1}, 2, []int{1, 2}},
		{"negative", []int{-1}, 0, []int{-1, 0}}, {"zero", []int{0}, 0, []int{0, 0}},
		{"many", []int{1, 2, 3}, 4, []int{1, 2, 3, 4}}, {"same", []int{7, 7}, 7, []int{7, 7, 7}},
		{"large", []int{100, 200}, 300, []int{100, 200, 300}}, {"extra cap", withCap([]int{1, 2}, 5), 3, []int{1, 2, 3}},
		{"mixed", []int{-1, 0, 1}, -2, []int{-1, 0, 1, -2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AppendOne(tc.items, tc.value); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAppendMany(t *testing.T) {
	cases := []struct {
		name                string
		items, values, want []int
	}{
		{"nil none", nil, nil, []int(nil)}, {"empty none", []int{}, []int{}, []int{}}, {"nil values", nil, []int{1, 2}, []int{1, 2}},
		{"one", []int{1}, []int{2}, []int{1, 2}}, {"many", []int{1}, []int{2, 3, 4}, []int{1, 2, 3, 4}},
		{"negative", []int{-1}, []int{0, 1}, []int{-1, 0, 1}}, {"zeros", []int{0}, []int{0, 0}, []int{0, 0, 0}},
		{"empty source", []int{}, []int{5, 6}, []int{5, 6}}, {"extra cap", withCap([]int{1, 2}, 6), []int{3, 4}, []int{1, 2, 3, 4}},
		{"mixed", []int{-2, -1}, []int{0, 1, 2}, []int{-2, -1, 0, 1, 2}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := AppendMany(tc.items, tc.values...); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCanAppendWithoutGrow(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		extra int
		want  bool
	}{
		{"nil zero", nil, 0, true}, {"nil one", nil, 1, false}, {"empty zero", []int{}, 0, true},
		{"full zero", []int{1, 2}, 0, true}, {"full one", []int{1, 2}, 1, false},
		{"space exact", withCap([]int{1, 2}, 5), 3, true}, {"space under", withCap([]int{1, 2}, 5), 2, true},
		{"space over", withCap([]int{1, 2}, 5), 4, false}, {"negative", withCap([]int{1}, 5), -1, false},
		{"empty cap ten", make([]int, 0, 10), 10, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanAppendWithoutGrow(tc.items, tc.extra); got != tc.want {
				t.Fatalf("len/cap=%d/%d extra=%d got %t want %t", len(tc.items), cap(tc.items), tc.extra, got, tc.want)
			}
		})
	}
}

func TestSliceKind(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  string
	}{
		{"nil", nil, "nil"}, {"empty literal", []int{}, "empty"}, {"empty make", make([]int, 0), "empty"},
		{"empty cap", make([]int, 0, 5), "empty"}, {"one zero", []int{0}, "filled"}, {"one", []int{1}, "filled"},
		{"many", []int{1, 2, 3}, "filled"}, {"negative", []int{-1}, "filled"}, {"subslice empty", []int{1, 2}[1:1], "empty"},
		{"subslice filled", []int{1, 2, 3}[1:3], "filled"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SliceKind(tc.items); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestAppendIndependent(t *testing.T) {
	cases := []struct {
		name         string
		backing      []int
		length       int
		values, want []int
	}{
		{"nil", nil, 0, []int{1}, []int{1}},
		{"empty backing", []int{}, 0, []int{1, 2}, []int{1, 2}},
		{"one no extra", []int{1}, 1, []int{2}, []int{1, 2}},
		{"shared free capacity", []int{1, 2, 90, 91, 92}, 2, []int{3, 4}, []int{1, 2, 3, 4}},
		{"overwrite sentinel risk", []int{5, 6, 70, 80}, 2, []int{7}, []int{5, 6, 7}},
		{"no values", []int{1, 2, 90}, 2, nil, []int{1, 2}},
		{"negative", []int{-1, -2, 99}, 2, []int{0, 1}, []int{-1, -2, 0, 1}},
		{"zeros", []int{0, 0, 9}, 2, []int{0}, []int{0, 0, 0}},
		{"full source", []int{1, 2, 3}, 3, []int{4}, []int{1, 2, 3, 4}},
		{"many", []int{10, 20, 90, 91, 92, 93}, 2, []int{30, 40, 50}, []int{10, 20, 30, 40, 50}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var source []int
			if tc.backing != nil {
				source = tc.backing[:tc.length]
			}
			var beforeBacking []int
			if tc.backing != nil {
				beforeBacking = append([]int{}, tc.backing...)
			}
			var beforeSource []int
			if source != nil {
				beforeSource = append([]int{}, source...)
			}
			got := AppendIndependent(source, tc.values...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if !reflect.DeepEqual(tc.backing, beforeBacking) {
				t.Fatalf("backing changed: %v, want %v", tc.backing, beforeBacking)
			}
			if !reflect.DeepEqual(source, beforeSource) {
				t.Fatalf("source changed: %v", source)
			}
			if len(got) > 0 && len(source) > 0 {
				got[0]++
				if source[0] != beforeSource[0] {
					t.Fatal("result shares source backing")
				}
			}
		})
	}
}

func withCap(values []int, capacity int) []int {
	items := make([]int, len(values), capacity)
	copy(items, values)
	return items
}

func TestExampleOutput(t *testing.T) {
	want := "first=10/true last=sql/true regular=2/3 limited=2/2 mutated=[10 999 30 40] appendSource=[10 20 30 777] appendPart=[20 30 777] limitedSource=[10 20 30 40] limitedPart=[20 30 777] grow=true kind=filled independent=[1 2 3 4] original=[1 2]"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
