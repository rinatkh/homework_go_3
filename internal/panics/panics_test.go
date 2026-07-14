package panics

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestSafeRun(t *testing.T) {
	cases := []struct {
		name string
		fn   func()
		want bool
	}{
		{"empty", func() {}, false}, {"assignment", func() { _ = 1 + 1 }, false}, {"panic string", func() { panic("boom") }, true},
		{"panic int", func() { panic(42) }, true}, {"panic error", func() { panic(errors.New("bad")) }, true},
		{"nested normal", func() { func() {}() }, false}, {"nested panic", func() { func() { panic("nested") }() }, true},
		{"defer normal", func() { defer func() {}() }, false}, {"panic nil pointer", func() { var p *int; _ = *p }, true},
		{"panic index", func() { _ = []int{1}[2] }, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SafeRun(tc.fn); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestRecoverMessage(t *testing.T) {
	cases := []struct {
		name string
		fn   func()
		want string
	}{
		{"normal", func() {}, ""}, {"string", func() { panic("boom") }, "boom"}, {"int", func() { panic(42) }, "42"},
		{"error", func() { panic(errors.New("bad")) }, "bad"}, {"empty string", func() { panic("") }, ""},
		{"unicode", func() { panic("ошибка") }, "ошибка"}, {"bool", func() { panic(true) }, "true"},
		{"slice text", func() { panic([]int{1, 2}) }, "[1 2]"}, {"nested", func() { func() { panic("nested") }() }, "nested"},
		{"nil pointer", func() { var p *int; _ = *p }, "runtime error: invalid memory address or nil pointer dereference"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RecoverMessage(tc.fn)
			if tc.name == "nil pointer" {
				if !strings.Contains(got, "invalid memory address") {
					t.Fatalf("got %q", got)
				}
				return
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestPanicIfEmpty(t *testing.T) {
	cases := []struct {
		name, text string
		wantPanic  bool
	}{
		{"empty", "", true}, {"space", " ", false}, {"go", "go", false}, {"unicode", "Привет", false},
		{"emoji", "🙂", false}, {"zero", "0", false}, {"newline", "\n", false}, {"tab", "\t", false},
		{"long", "long text", false}, {"two spaces", "  ", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			panicked, value := capturePanic(func() { PanicIfEmpty(tc.text) })
			if panicked != tc.wantPanic {
				t.Fatalf("panicked=%t, want %t", panicked, tc.wantPanic)
			}
			if tc.wantPanic && value != "empty text" {
				t.Fatalf("panic=%v", value)
			}
		})
	}
}

func TestMustPositive(t *testing.T) {
	cases := []struct {
		name      string
		n, want   int
		wantPanic bool
	}{
		{"one", 1, 1, false}, {"two", 2, 2, false}, {"large", 1000, 1000, false}, {"max sample", 999999, 999999, false},
		{"zero", 0, 0, true}, {"minus one", -1, 0, true}, {"negative", -10, 0, true}, {"large negative", -9999, 0, true},
		{"forty two", 42, 42, false}, {"seven", 7, 7, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			panicked, value := capturePanic(func() { got = MustPositive(tc.n) })
			if panicked != tc.wantPanic {
				t.Fatalf("panicked=%t, want %t", panicked, tc.wantPanic)
			}
			if tc.wantPanic {
				if value != "not positive" {
					t.Fatalf("panic=%v", value)
				}
				return
			}
			if got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestParseOrPanic(t *testing.T) {
	cases := []struct {
		name, raw string
		want      int
		wantPanic bool
	}{
		{"positive", "42", 42, false}, {"zero", "0", 0, false}, {"negative", "-7", -7, false},
		{"plus", "+5", 5, false}, {"leading zeros", "007", 7, false}, {"letters", "abc", 0, true},
		{"empty", "", 0, true}, {"spaces", " 2 ", 0, true}, {"decimal", "2.5", 0, true}, {"mixed", "1x", 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			panicked, _ := capturePanic(func() { got = ParseOrPanic(tc.raw) })
			if panicked != tc.wantPanic {
				t.Fatalf("panicked=%t, want %t", panicked, tc.wantPanic)
			}
			if !tc.wantPanic && got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestTryParse(t *testing.T) {
	cases := []struct {
		name, raw string
		want      int
		ok        bool
	}{
		{"positive", "42", 42, true}, {"zero", "0", 0, true}, {"negative", "-7", -7, true}, {"plus", "+5", 5, true},
		{"leading zeros", "007", 7, true}, {"letters", "abc", 0, false}, {"empty", "", 0, false},
		{"spaces", " 2 ", 0, false}, {"decimal", "2.5", 0, false}, {"mixed", "1x", 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := TryParse(tc.raw)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %d/%t, want %d/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestSafeIndex(t *testing.T) {
	cases := []struct {
		name        string
		items       []int
		index, want int
		ok          bool
	}{
		{"first", []int{10, 20, 30}, 0, 10, true}, {"middle", []int{10, 20, 30}, 1, 20, true},
		{"last", []int{10, 20, 30}, 2, 30, true}, {"negative index", []int{1, 2}, -1, 0, false},
		{"equal len", []int{1, 2}, 2, 0, false}, {"far", []int{1}, 10, 0, false}, {"nil", nil, 0, 0, false},
		{"empty", []int{}, 0, 0, false}, {"zero value", []int{0}, 0, 0, true}, {"negative value", []int{-5}, 0, -5, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := SafeIndex(tc.items, tc.index)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %d/%t, want %d/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestMustGet(t *testing.T) {
	cases := []struct {
		name      string
		items     []string
		index     int
		want      string
		wantPanic bool
	}{
		{"first", []string{"a", "b"}, 0, "a", false}, {"last", []string{"a", "b"}, 1, "b", false},
		{"empty value", []string{""}, 0, "", false}, {"unicode", []string{"Привет"}, 0, "Привет", false},
		{"emoji", []string{"🙂"}, 0, "🙂", false}, {"negative index", []string{"a"}, -1, "", true},
		{"equal len", []string{"a"}, 1, "", true}, {"nil", nil, 0, "", true}, {"empty slice", []string{}, 0, "", true},
		{"far", []string{"a", "b"}, 10, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got string
			panicked, _ := capturePanic(func() { got = MustGet(tc.items, tc.index) })
			if panicked != tc.wantPanic {
				t.Fatalf("panicked=%t, want %t", panicked, tc.wantPanic)
			}
			if !tc.wantPanic && got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestSafeMustGet(t *testing.T) {
	cases := []struct {
		name  string
		items []string
		index int
		want  string
		ok    bool
	}{
		{"first", []string{"a", "b"}, 0, "a", true}, {"last", []string{"a", "b"}, 1, "b", true},
		{"empty value", []string{""}, 0, "", true}, {"unicode", []string{"Привет"}, 0, "Привет", true},
		{"emoji", []string{"🙂"}, 0, "🙂", true}, {"negative index", []string{"a"}, -1, "", false},
		{"equal len", []string{"a"}, 1, "", false}, {"nil", nil, 0, "", false}, {"empty slice", []string{}, 0, "", false},
		{"far", []string{"a", "b"}, 10, "", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := SafeMustGet(tc.items, tc.index)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %q/%t, want %q/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestPanicToError(t *testing.T) {
	cases := []struct {
		name string
		fn   func()
		want string
	}{
		{"normal", func() {}, ""}, {"string", func() { panic("boom") }, "panic: boom"},
		{"int", func() { panic(42) }, "panic: 42"}, {"error", func() { panic(errors.New("bad")) }, "panic: bad"},
		{"empty", func() { panic("") }, "panic: "}, {"unicode", func() { panic("ошибка") }, "panic: ошибка"},
		{"bool", func() { panic(true) }, "panic: true"}, {"nested", func() { func() { panic("nested") }() }, "panic: nested"},
		{"index", func() { _ = []int{1}[2] }, "panic: runtime error: index out of range [2] with length 1"},
		{"normal defer", func() { defer func() {}() }, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := PanicToError(tc.fn)
			if tc.want == "" {
				if err != nil {
					t.Fatalf("got %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatal("got nil")
			}
			if tc.name == "index" {
				if !strings.Contains(err.Error(), "panic: runtime error: index out of range") {
					t.Fatalf("got %q", err.Error())
				}
				return
			}
			if err.Error() != tc.want {
				t.Fatalf("got %q, want %q", err.Error(), tc.want)
			}
		})
	}
}

func TestDeferBeforePanic(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			got := DeferBeforePanic()
			want := []string{"body", "defer"}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestRecoverOutsideDefer(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(caseName(i), func(t *testing.T) {
			if RecoverOutsideDefer() {
				t.Fatal("recover outside defer must return nil")
			}
		})
	}
}

func TestRunSequence(t *testing.T) {
	cases := []struct {
		name                          string
		panicAt, total, wantCompleted int
		wantPanic                     bool
	}{
		{"none empty", -1, 0, 0, false}, {"none one", -1, 1, 1, false}, {"none three", -1, 3, 3, false},
		{"none five", -1, 5, 5, false}, {"panic first", 0, 3, 0, true}, {"panic second", 1, 3, 1, true},
		{"panic last", 2, 3, 2, true}, {"panic middle five", 2, 5, 2, true}, {"panic only", 0, 1, 0, true},
		{"none ten", -1, 10, 10, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			executed := 0
			functions := make([]func(), tc.total)
			for i := 0; i < tc.total; i++ {
				index := i
				functions[i] = func() {
					if index == tc.panicAt {
						panic("stop")
					}
					executed++
				}
			}
			var got int
			panicked, _ := capturePanic(func() { got = RunSequence(functions) })
			if panicked != tc.wantPanic {
				t.Fatalf("panicked=%t, want %t", panicked, tc.wantPanic)
			}
			if executed != tc.wantCompleted {
				t.Fatalf("executed=%d, want %d", executed, tc.wantCompleted)
			}
			if !tc.wantPanic && got != tc.wantCompleted {
				t.Fatalf("got=%d, want %d", got, tc.wantCompleted)
			}
		})
	}
}

func TestRunSequenceSafe(t *testing.T) {
	cases := []struct {
		name                string
		pattern             []bool
		completed, panicked int
	}{
		{"empty", nil, 0, 0}, {"one normal", []bool{false}, 1, 0}, {"one panic", []bool{true}, 0, 1},
		{"mixed", []bool{false, true, false}, 2, 1}, {"all normal", []bool{false, false, false}, 3, 0},
		{"all panic", []bool{true, true, true}, 0, 3}, {"starts panic", []bool{true, false, false}, 2, 1},
		{"ends panic", []bool{false, false, true}, 2, 1}, {"alternating", []bool{true, false, true, false}, 2, 2},
		{"five", []bool{false, true, true, false, true}, 2, 3},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			functions := make([]func(), len(tc.pattern))
			for i, shouldPanic := range tc.pattern {
				panicFlag := shouldPanic
				functions[i] = func() {
					calls++
					if panicFlag {
						panic("x")
					}
				}
			}
			completed, panicked := RunSequenceSafe(functions)
			if completed != tc.completed || panicked != tc.panicked {
				t.Fatalf("got %d/%d, want %d/%d", completed, panicked, tc.completed, tc.panicked)
			}
			if calls != len(tc.pattern) {
				t.Fatalf("calls=%d, want %d", calls, len(tc.pattern))
			}
		})
	}
}

func TestMustNonNil(t *testing.T) {
	values := []*int{intPointer(0), intPointer(1), intPointer(-1), nil, intPointer(42), intPointer(1000), nil, intPointer(-999), intPointer(7), intPointer(9)}
	for i, value := range values {
		t.Run(caseName(i), func(t *testing.T) {
			var got int
			panicked, panicValue := capturePanic(func() { got = MustNonNil(value) })
			if value == nil {
				if !panicked || panicValue != "nil pointer" {
					t.Fatalf("panicked=%t value=%v", panicked, panicValue)
				}
				return
			}
			if panicked || got != *value {
				t.Fatalf("got %d panicked=%t, want %d", got, panicked, *value)
			}
		})
	}
}

func capturePanic(fn func()) (panicked bool, value any) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			value = r
		}
	}()
	fn()
	return false, nil
}

func intPointer(value int) *int { return &value }

func caseName(i int) string { const digits = "0123456789"; return "case_" + string(digits[i]) }

func TestExampleOutput(t *testing.T) {
	want := "normal=false boom=true message=broken parse=42/true invalid=0/false index=20/true missing=0/false error=panic: bad state events=[body defer] outside=false sequence=2/1 pointer=7"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
