package functions

import (
	"errors"
	"reflect"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want int
		wantErr    bool
	}{
		{"exact", 10, 2, 5, false}, {"truncate", 7, 2, 3, false}, {"zero numerator", 0, 5, 0, false},
		{"negative numerator", -10, 2, -5, false}, {"negative denominator", 10, -2, -5, false},
		{"both negative", -10, -2, 5, false}, {"larger denominator", 3, 10, 0, false},
		{"divide by one", 99, 1, 99, false}, {"divide by zero", 10, 0, 0, true}, {"zero by zero", 0, 0, 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SafeDivide(tc.a, tc.b)
			if got != tc.want || (err != nil) != tc.wantErr {
				t.Fatalf("got %d/%v, want %d err=%t", got, err, tc.want, tc.wantErr)
			}
			if tc.wantErr && err.Error() != "division by zero" {
				t.Fatalf("error=%q", err.Error())
			}
		})
	}
}

func TestFindUserByID(t *testing.T) {
	users := []User{{ID: 1, Name: "Maria", Active: true}, {ID: 2, Name: "Alex"}, {ID: 3, Name: "Ira", Active: true}, {ID: -1, Name: "Temp"}}
	cases := []struct {
		name  string
		users []User
		id    int
		want  User
		ok    bool
	}{
		{"nil", nil, 1, User{}, false}, {"empty", []User{}, 1, User{}, false}, {"first", users, 1, users[0], true},
		{"middle", users, 2, users[1], true}, {"last positive", users, 3, users[2], true}, {"negative id", users, -1, users[3], true},
		{"missing", users, 10, User{}, false}, {"zero missing", users, 0, User{}, false},
		{"duplicate returns first", []User{{ID: 5, Name: "first"}, {ID: 5, Name: "second"}}, 5, User{ID: 5, Name: "first"}, true},
		{"inactive still found", []User{{ID: 7, Name: "A", Active: false}}, 7, User{ID: 7, Name: "A"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := FindUserByID(tc.users, tc.id)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %+v/%t, want %+v/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestFindActiveUser(t *testing.T) {
	users := []User{{ID: 1, Name: "Maria", Active: true}, {ID: 2, Name: "Alex"}, {ID: 3, Name: "Ira", Active: true}, {ID: -1, Name: "Temp", Active: true}}
	cases := []struct {
		name  string
		users []User
		id    int
		want  User
		ok    bool
	}{
		{"nil", nil, 1, User{}, false}, {"empty", []User{}, 1, User{}, false}, {"active first", users, 1, users[0], true},
		{"inactive", users, 2, User{}, false}, {"active middle", users, 3, users[2], true}, {"active negative id", users, -1, users[3], true},
		{"missing", users, 10, User{}, false}, {"zero missing", users, 0, User{}, false},
		{"duplicate inactive then active", []User{{ID: 5, Name: "off"}, {ID: 5, Name: "on", Active: true}}, 5, User{ID: 5, Name: "on", Active: true}, true},
		{"duplicate active first", []User{{ID: 7, Name: "first", Active: true}, {ID: 7, Name: "second", Active: true}}, 7, User{ID: 7, Name: "first", Active: true}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := FindActiveUser(tc.users, tc.id)
			if got != tc.want || ok != tc.ok {
				t.Fatalf("got %+v/%t, want %+v/%t", got, ok, tc.want, tc.ok)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	cases := []struct {
		name  string
		items []int
		want  int
	}{
		{"none", nil, 0}, {"one", []int{7}, 7}, {"two", []int{1, 2}, 3}, {"four", []int{1, 2, 3, 4}, 10},
		{"negative", []int{-1, -2, -3}, -6}, {"mixed", []int{-5, 0, 5}, 0}, {"zeros", []int{0, 0, 0}, 0},
		{"large", []int{100, 200, 300}, 600}, {"cancel", []int{10, -3, -7}, 0}, {"many", []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 10},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SumAll(tc.items...); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want int
		op         func(int, int) int
	}{
		{"add", 2, 3, 5, func(a, b int) int { return a + b }}, {"subtract", 5, 3, 2, func(a, b int) int { return a - b }},
		{"multiply", 4, 6, 24, func(a, b int) int { return a * b }}, {"max", 4, 9, 9, func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}},
		{"min", 4, 9, 4, func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}}, {"first", 7, 8, 7, func(a, b int) int { return a }},
		{"second", 7, 8, 8, func(a, b int) int { return b }}, {"negative", -2, 3, -6, func(a, b int) int { return a * b }},
		{"zero", 0, 5, 5, func(a, b int) int { return a + b }}, {"difference absolute", 3, 10, 7, func(a, b int) int {
			if a > b {
				return a - b
			}
			return b - a
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Apply(tc.a, tc.b, tc.op); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestApplyIf(t *testing.T) {
	cases := []struct {
		name      string
		value     int
		allowed   bool
		want      int
		transform func(int) int
		wantCalls int
	}{
		{"allowed double", 3, true, 6, func(v int) int { return v * 2 }, 1}, {"blocked double", 3, false, 3, func(v int) int { return v * 2 }, 0},
		{"allowed zero", 0, true, 1, func(v int) int { return v + 1 }, 1}, {"blocked zero", 0, false, 0, func(v int) int { return v + 1 }, 0},
		{"negative allowed", -2, true, 2, func(v int) int { return -v }, 1}, {"negative blocked", -2, false, -2, func(v int) int { return -v }, 0},
		{"square", 4, true, 16, func(v int) int { return v * v }, 1}, {"identity", 7, true, 7, func(v int) int { return v }, 1},
		{"large blocked", 1000, false, 1000, func(v int) int { return 0 }, 0}, {"subtract", 10, true, 7, func(v int) int { return v - 3 }, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			got := ApplyIf(tc.value, tc.allowed, func(v int) int { calls++; return tc.transform(v) })
			if got != tc.want || calls != tc.wantCalls {
				t.Fatalf("got %d calls=%d, want %d/%d", got, calls, tc.want, tc.wantCalls)
			}
		})
	}
}

func TestNewCounter(t *testing.T) {
	cases := []struct {
		name  string
		start int
		want  []int
	}{
		{"zero", 0, []int{1, 2, 3}}, {"five", 5, []int{6, 7, 8}}, {"negative", -2, []int{-1, 0, 1}},
		{"one", 1, []int{2, 3, 4}}, {"large", 100, []int{101, 102, 103}}, {"minus one", -1, []int{0, 1, 2}},
		{"ten", 10, []int{11, 12, 13}}, {"forty two", 42, []int{43, 44, 45}}, {"minus ten", -10, []int{-9, -8, -7}},
		{"thousand", 1000, []int{1001, 1002, 1003}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			counter := NewCounter(tc.start)
			got := []int{counter(), counter(), counter()}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			other := NewCounter(tc.start)
			if other() != tc.start+1 {
				t.Fatal("counters share state")
			}
		})
	}
}

func TestNewAccumulator(t *testing.T) {
	cases := []struct {
		name         string
		initial      int
		values, want []int
	}{
		{"zero", 0, []int{1, 2, 3}, []int{1, 3, 6}}, {"ten", 10, []int{5, -2, 1}, []int{15, 13, 14}},
		{"negative initial", -5, []int{2, 2, 2}, []int{-3, -1, 1}}, {"zeros", 7, []int{0, 0, 0}, []int{7, 7, 7}},
		{"one call", 1, []int{9}, []int{10}}, {"negative values", 0, []int{-1, -2, -3}, []int{-1, -3, -6}},
		{"mixed", 100, []int{-50, 25, -10}, []int{50, 75, 65}}, {"large", 1000, []int{1000, 1000}, []int{2000, 3000}},
		{"cancel", 5, []int{10, -10}, []int{15, 5}}, {"four", 0, []int{1, 1, 1, 1}, []int{1, 2, 3, 4}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			acc := NewAccumulator(tc.initial)
			got := make([]int, 0, len(tc.values))
			for _, value := range tc.values {
				got = append(got, acc(value))
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMakeMultiplier(t *testing.T) {
	cases := []struct {
		name                string
		factor, value, want int
	}{
		{"double", 2, 5, 10}, {"triple", 3, 7, 21}, {"zero factor", 0, 99, 0}, {"zero value", 10, 0, 0},
		{"negative factor", -2, 4, -8}, {"negative value", 3, -5, -15}, {"both negative", -3, -5, 15},
		{"one", 1, 42, 42}, {"large", 100, 20, 2000}, {"minus one", -1, 9, -9},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			multiply := MakeMultiplier(tc.factor)
			if got := multiply(tc.value); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
			if got := multiply(tc.value + 1); got != (tc.value+1)*tc.factor {
				t.Fatal("factor was not captured")
			}
		})
	}
}

func TestDeferOrder(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(testCaseName(i), func(t *testing.T) {
			if got := DeferOrder(); got != "body-second-first" {
				t.Fatalf("got %q", got)
			}
		})
	}
}

func TestCaptureDeferArgument(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(testCaseName(i), func(t *testing.T) {
			if got := CaptureDeferArgument(); got != "first" {
				t.Fatalf("got %q, want first", got)
			}
		})
	}
}

func TestReadDeferredVariable(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(testCaseName(i), func(t *testing.T) {
			if got := ReadDeferredVariable(); got != "second" {
				t.Fatalf("got %q, want second", got)
			}
		})
	}
}

func TestIncrementNamedResult(t *testing.T) {
	values := []int{0, 1, -1, 5, 10, -10, 99, 1000, -999, 42}
	for i, value := range values {
		t.Run(testCaseName(i), func(t *testing.T) {
			if got := IncrementNamedResult(value); got != value+1 {
				t.Fatalf("got %d, want %d", got, value+1)
			}
		})
	}
}

func TestRunWithCleanup(t *testing.T) {
	cases := []struct{ name, action, cleanup string }{
		{"basic", "action", "cleanup"}, {"open close", "open", "close"}, {"lock unlock", "lock", "unlock"},
		{"start stop", "start", "stop"}, {"empty action", "", "cleanup"}, {"empty cleanup", "action", ""},
		{"both empty", "", ""}, {"unicode", "работа", "очистка"}, {"emoji", "▶", "■"}, {"long", "do something", "release resource"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			calls := []string{}
			got := RunWithCleanup(func() string { calls = append(calls, "action"); return tc.action }, func() string { calls = append(calls, "cleanup"); return tc.cleanup })
			want := []string{tc.action, tc.cleanup}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
			if !reflect.DeepEqual(calls, []string{"action", "cleanup"}) {
				t.Fatalf("call order=%v", calls)
			}
		})
	}
}

func TestChooseOperation(t *testing.T) {
	cases := []struct {
		name, operation string
		a, b, want      int
		ok              bool
	}{
		{"add", "add", 2, 3, 5, true}, {"sub", "sub", 10, 4, 6, true}, {"mul", "mul", 6, 7, 42, true},
		{"unknown", "div", 8, 2, 0, false}, {"empty", "", 1, 2, 0, false}, {"case sensitive", "ADD", 1, 2, 0, false},
		{"add negative", "add", -2, 5, 3, true}, {"sub negative", "sub", -2, -5, 3, true},
		{"mul zero", "mul", 10, 0, 0, true}, {"spaces", " add ", 1, 2, 0, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			op, ok := ChooseOperation(tc.operation)
			if ok != tc.ok {
				t.Fatalf("ok=%t, want %t", ok, tc.ok)
			}
			if !tc.ok {
				if op != nil {
					t.Fatal("unknown operation returned function")
				}
				return
			}
			if op == nil {
				t.Fatal("nil function")
			}
			if got := op(tc.a, tc.b); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func testCaseName(i int) string {
	const digits = "0123456789"
	return "case_" + string(digits[i])
}

var _ = errors.New

func TestExampleOutput(t *testing.T) {
	want := "divide=5/false user=Maria/true sum=10 apply=13 allowed=16 counter=6,7 accumulator=15,13 multiply=21 defer=body-second-first captured=first deferred=second named=10 cleanup=[action cleanup] operation=20/true"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
