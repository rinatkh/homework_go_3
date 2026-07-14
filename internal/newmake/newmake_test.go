package newmake

import (
	"reflect"
	"testing"
)

func TestNewInt(t *testing.T) {
	values := []int{0, 1, -1, 7, 26, 100, -100, 9999, -9999, 42}
	for i, value := range values {
		t.Run(caseName(i), func(t *testing.T) {
			p := NewInt(value)
			if p == nil {
				t.Fatal("nil pointer")
			}
			if *p != value {
				t.Fatalf("got %d, want %d", *p, value)
			}
		})
	}
}

func TestNewString(t *testing.T) {
	values := []string{"", "go", "Maria", "Привет", "🙂", " ", "0", "long text", "a@b", "last"}
	for i, value := range values {
		t.Run(caseName(i), func(t *testing.T) {
			p := NewString(value)
			if p == nil {
				t.Fatal("nil pointer")
			}
			if *p != value {
				t.Fatalf("got %q, want %q", *p, value)
			}
		})
	}
}

func TestNewZeroUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			p := NewZeroUser()
			if p == nil {
				t.Fatal("nil pointer")
			}
			if p.Name != "" || p.Scores != nil {
				t.Fatalf("got %+v", *p)
			}
		})
	}
}

func TestNewNamedUser(t *testing.T) {
	names := []string{"Maria", "", "Alex", "Мария", "🙂", " ", "A", "Long Name", "0", "Go"}
	for i, name := range names {
		t.Run(caseName(i), func(t *testing.T) {
			p := NewNamedUser(name)
			if p == nil {
				t.Fatal("nil pointer")
			}
			if p.Name != name {
				t.Fatalf("name=%q, want %q", p.Name, name)
			}
			if p.Scores != nil {
				t.Fatalf("Scores=%v, want nil", p.Scores)
			}
		})
	}
}

func TestMakeInts(t *testing.T) {
	lengths := []int{0, 1, 2, 3, 4, 5, 8, 10, 16, 20}
	for i, length := range lengths {
		t.Run(caseName(i), func(t *testing.T) {
			items := MakeInts(length)
			if items == nil {
				t.Fatal("nil slice")
			}
			if len(items) != length {
				t.Fatalf("len=%d, want %d", len(items), length)
			}
			for index, value := range items {
				if value != 0 {
					t.Fatalf("items[%d]=%d, want 0", index, value)
				}
			}
		})
	}
}

func TestMakeIntsWithCapacity(t *testing.T) {
	cases := []struct{ length, capacity int }{{0, 0}, {0, 1}, {1, 1}, {0, 5}, {2, 5}, {5, 5}, {3, 10}, {1, 8}, {8, 16}, {10, 20}}
	for i, tc := range cases {
		t.Run(caseName(i), func(t *testing.T) {
			items := MakeIntsWithCapacity(tc.length, tc.capacity)
			if items == nil {
				t.Fatal("nil slice")
			}
			if len(items) != tc.length || cap(items) != tc.capacity {
				t.Fatalf("len/cap=%d/%d, want %d/%d", len(items), cap(items), tc.length, tc.capacity)
			}
		})
	}
}

func TestMakeAndAppend(t *testing.T) {
	cases := []struct {
		capacity     int
		values, want []int
	}{
		{0, nil, []int{}}, {3, []int{1, 2, 3}, []int{1, 2, 3}}, {10, []int{1}, []int{1}},
		{1, []int{1, 2}, []int{1, 2}}, {0, []int{-1, 0, 1}, []int{-1, 0, 1}}, {5, []int{7, 7, 7}, []int{7, 7, 7}},
		{2, []int{100, 200, 300}, []int{100, 200, 300}}, {4, []int{}, []int{}}, {1, []int{42}, []int{42}},
		{8, []int{9, 8, 7, 6}, []int{9, 8, 7, 6}},
	}
	for i, tc := range cases {
		t.Run(caseName(i), func(t *testing.T) {
			got := MakeAndAppend(tc.capacity, tc.values...)
			if got == nil {
				t.Fatal("nil slice")
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			if len(tc.values) <= tc.capacity && cap(got) < tc.capacity {
				t.Fatalf("cap=%d, expected at least %d", cap(got), tc.capacity)
			}
		})
	}
}

func TestMakeEmptyStrings(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			items := MakeEmptyStrings()
			if items == nil {
				t.Fatal("nil slice")
			}
			if len(items) != 0 {
				t.Fatalf("len=%d", len(items))
			}
			items = append(items, "go")
			if !reflect.DeepEqual(items, []string{"go"}) {
				t.Fatalf("append failed: %v", items)
			}
		})
	}
}

func TestSliceState(t *testing.T) {
	makeWithCap := func(values []int, capacity int) []int {
		items := make([]int, len(values), capacity)
		copy(items, values)
		return items
	}
	cases := []struct {
		name             string
		items            []int
		wantLen, wantCap int
		wantNil          bool
	}{
		{"nil", nil, 0, 0, true}, {"empty literal", []int{}, 0, 0, false}, {"one", []int{1}, 1, 1, false},
		{"three", []int{1, 2, 3}, 3, 3, false}, {"empty cap five", make([]int, 0, 5), 0, 5, false},
		{"len two cap five", makeWithCap([]int{1, 2}, 5), 2, 5, false}, {"len five cap ten", makeWithCap([]int{1, 2, 3, 4, 5}, 10), 5, 10, false},
		{"zero len cap one", make([]int, 0, 1), 0, 1, false}, {"negative values", []int{-1, -2}, 2, 2, false},
		{"subslice", makeWithCap([]int{10, 20, 30, 40}, 8)[1:3], 2, 7, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			length, capacity, isNil := SliceState(tc.items)
			if length != tc.wantLen || capacity != tc.wantCap || isNil != tc.wantNil {
				t.Fatalf("got %d/%d/%t, want %d/%d/%t", length, capacity, isNil, tc.wantLen, tc.wantCap, tc.wantNil)
			}
		})
	}
}

func TestEmptySliceState(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			length, capacity, isNil := EmptySliceState()
			if length != 0 || capacity != 0 || isNil {
				t.Fatalf("got %d/%d/%t, want 0/0/false", length, capacity, isNil)
			}
		})
	}
}

func TestMakeUserScores(t *testing.T) {
	cases := []struct {
		name     string
		userName string
		count    int
	}{{"zero", "Maria", 0}, {"one", "Alex", 1}, {"three", "Ira", 3}, {"empty name", "", 2}, {"unicode", "Мария", 4}, {"emoji", "🙂", 1}, {"space", " ", 5}, {"large", "Student", 10}, {"zero empty", "", 0}, {"another", "Go", 7}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			user := MakeUserScores(tc.userName, tc.count)
			if user == nil {
				t.Fatal("nil user")
			}
			if user.Name != tc.userName {
				t.Fatalf("name=%q, want %q", user.Name, tc.userName)
			}
			if user.Scores == nil {
				t.Fatal("nil Scores")
			}
			if len(user.Scores) != tc.count {
				t.Fatalf("len=%d, want %d", len(user.Scores), tc.count)
			}
			for _, score := range user.Scores {
				if score != 0 {
					t.Fatalf("non-zero score %d", score)
				}
			}
		})
	}
}

func TestNewSlicePointer(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			p := NewSlicePointer()
			if p == nil {
				t.Fatal("nil pointer")
			}
			if *p != nil {
				t.Fatalf("slice=%v, want nil", *p)
			}
			if len(*p) != 0 || cap(*p) != 0 {
				t.Fatalf("len/cap=%d/%d", len(*p), cap(*p))
			}
		})
	}
}

func TestMakeSlicePointer(t *testing.T) {
	capacities := []int{0, 1, 2, 3, 4, 5, 8, 10, 16, 20}
	for i, capacity := range capacities {
		t.Run(caseName(i), func(t *testing.T) {
			p := MakeSlicePointer(capacity)
			if p == nil {
				t.Fatal("nil pointer")
			}
			if *p == nil {
				t.Fatal("nil slice")
			}
			if len(*p) != 0 || cap(*p) != capacity {
				t.Fatalf("len/cap=%d/%d, want 0/%d", len(*p), cap(*p), capacity)
			}
			*p = append(*p, 1)
			if len(*p) != 1 || (*p)[0] != 1 {
				t.Fatalf("append failed: %v", *p)
			}
		})
	}
}

func TestMakeStringIntMap(t *testing.T) {
	keys := []string{"go", "sql", "redis", "", "Привет", "🙂", "a", "b", "long key", "last"}
	for i, key := range keys {
		t.Run(caseName(i), func(t *testing.T) {
			values := MakeStringIntMap()
			if values == nil {
				t.Fatal("nil map")
			}
			if len(values) != 0 {
				t.Fatalf("len=%d, want 0", len(values))
			}
			values[key] = i
			if values[key] != i {
				t.Fatalf("value=%d, want %d", values[key], i)
			}
			another := MakeStringIntMap()
			if len(another) != 0 {
				t.Fatal("calls share map state")
			}
		})
	}
}

func TestMakeBoolChannel(t *testing.T) {
	buffers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, buffer := range buffers {
		t.Run(caseName(i), func(t *testing.T) {
			ch := MakeBoolChannel(buffer)
			if ch == nil {
				t.Fatal("nil channel")
			}
			defer close(ch)
			if cap(ch) != buffer {
				t.Fatalf("cap=%d, want %d", cap(ch), buffer)
			}
			if buffer > 0 {
				ch <- true
				if got := <-ch; !got {
					t.Fatal("got false")
				}
			}
		})
	}
}

func caseName(i int) string {
	const digits = "0123456789"
	return "case_" + string(digits[i])
}

func TestExampleOutput(t *testing.T) {
	want := "int=26 string=go slice=[0 0] len=2 cap=5 nil=false appended=[10 20 30] emptyNil=false user=Maria/3 newSliceNil=true mapNil=false chanCap=2"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
