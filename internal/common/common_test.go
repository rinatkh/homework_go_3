package common

import (
	"errors"
	"reflect"
	"testing"
)

func TestActiveUserReport(t *testing.T) {
	cases := []struct {
		name    string
		users   []User
		want    string
		wantErr bool
	}{
		{"nil", nil, "", true}, {"empty", []User{}, "", true}, {"one active", []User{{Name: "Maria", Active: true}}, "active=1 names=Maria", false},
		{"one inactive", []User{{Name: "Alex"}}, "", true}, {"mixed", []User{{Name: "Maria", Active: true}, {Name: "Alex"}, {Name: "Ira", Active: true}}, "active=2 names=Maria,Ira", false},
		{"all active", []User{{Name: "A", Active: true}, {Name: "B", Active: true}, {Name: "C", Active: true}}, "active=3 names=A,B,C", false},
		{"empty active name", []User{{Name: "", Active: true}}, "active=1 names=", false},
		{"unicode", []User{{Name: "Мария", Active: true}, {Name: "Ира", Active: true}}, "active=2 names=Мария,Ира", false},
		{"order", []User{{Name: "3", Active: true}, {Name: "1", Active: true}, {Name: "2", Active: true}}, "active=3 names=3,1,2", false},
		{"inactive around", []User{{Name: "off"}, {Name: "on", Active: true}, {Name: "off2"}}, "active=1 names=on", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ActiveUserReport(tc.users)
			if got != tc.want || (err != nil) != tc.wantErr {
				t.Fatalf("got %q/%v, want %q err=%t", got, err, tc.want, tc.wantErr)
			}
			if tc.wantErr && !errors.Is(err, ErrNoActiveUsers) {
				t.Fatalf("got %v, want ErrNoActiveUsers", err)
			}
		})
	}
}

func TestDiscountOrder(t *testing.T) {
	cases := []struct {
		name      string
		order     Order
		percent   int
		want      Order
		wantField string
	}{
		{"no discount", Order{ID: 1, Amount: 1000}, 0, Order{ID: 1, Amount: 1000}, ""},
		{"ten", Order{ID: 1, Amount: 1000}, 10, Order{ID: 1, Amount: 900}, ""},
		{"full", Order{ID: 2, Amount: 500, Paid: true}, 100, Order{ID: 2, Amount: 0, Paid: true}, ""},
		{"twenty five", Order{ID: 3, Amount: 200}, 25, Order{ID: 3, Amount: 150}, ""},
		{"truncate", Order{ID: 4, Amount: 99}, 10, Order{ID: 4, Amount: 89}, ""},
		{"negative amount", Order{ID: 5, Amount: -1}, 10, Order{}, "amount"},
		{"negative percent", Order{ID: 6, Amount: 100}, -1, Order{}, "percent"},
		{"percent over", Order{ID: 7, Amount: 100}, 101, Order{}, "percent"},
		{"zero amount", Order{ID: 8, Amount: 0}, 50, Order{ID: 8, Amount: 0}, ""},
		{"paid preserved", Order{ID: 9, Amount: 333, Paid: true}, 33, Order{ID: 9, Amount: 223, Paid: true}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			original := tc.order
			got, err := DiscountOrder(tc.order, tc.percent)
			if tc.order != original {
				t.Fatal("input changed")
			}
			if tc.wantField == "" {
				if err != nil || got != tc.want {
					t.Fatalf("got %+v/%v, want %+v/nil", got, err, tc.want)
				}
				return
			}
			if err == nil {
				t.Fatal("got nil error")
			}
			var fieldErr *FieldError
			if !errors.As(err, &fieldErr) {
				t.Fatalf("expected FieldError, got %T", err)
			}
			if fieldErr.Field != tc.wantField {
				t.Fatalf("field=%q, want %q", fieldErr.Field, tc.wantField)
			}
		})
	}
}

func TestAppendToIsolatedWindow(t *testing.T) {
	cases := []struct {
		name                   string
		items                  []int
		low, high              int
		values                 []int
		wantSource, wantWindow []int
		wantErr                bool
	}{
		{"middle", []int{10, 20, 30, 40}, 1, 3, []int{777}, []int{10, 20, 30, 40}, []int{20, 30, 777}, false},
		{"start", []int{1, 2, 3}, 0, 2, []int{8}, []int{1, 2, 3}, []int{1, 2, 8}, false},
		{"end", []int{1, 2, 3}, 1, 3, []int{4, 5}, []int{1, 2, 3}, []int{2, 3, 4, 5}, false},
		{"empty window", []int{1, 2, 3}, 1, 1, []int{9}, []int{1, 2, 3}, []int{9}, false},
		{"no values", []int{1, 2, 3}, 0, 2, nil, []int{1, 2, 3}, []int{1, 2}, false},
		{"nil valid", nil, 0, 0, []int{1}, nil, []int{1}, false},
		{"low negative", []int{1, 2}, -1, 1, []int{3}, nil, nil, true},
		{"high below low", []int{1, 2}, 2, 1, []int{3}, nil, nil, true},
		{"high too large", []int{1, 2}, 0, 3, []int{3}, nil, nil, true},
		{"negative values", []int{-1, -2, -3}, 1, 2, []int{0, 1}, []int{-1, -2, -3}, []int{-2, 0, 1}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			before := append([]int(nil), tc.items...)
			source, window, err := AppendToIsolatedWindow(tc.items, tc.low, tc.high, tc.values...)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v, wantErr=%t", err, tc.wantErr)
			}
			if !reflect.DeepEqual(source, tc.wantSource) || !reflect.DeepEqual(window, tc.wantWindow) {
				t.Fatalf("source=%v window=%v, want %v %v", source, window, tc.wantSource, tc.wantWindow)
			}
			if !tc.wantErr && !reflect.DeepEqual(tc.items, before) {
				t.Fatalf("input changed: %v, want %v", tc.items, before)
			}
			if len(tc.values) > 0 && len(window) > 0 && len(tc.items) > 0 {
				window[0]++
				if len(before) > 0 && tc.items[0] != before[0] {
					t.Fatal("window unexpectedly changes first source element after append")
				}
			}
		})
	}
}

func TestExampleOutput(t *testing.T) {
	want := `report="active=2 names=Maria,Ira"/false order={ID:7 Amount:1000 Paid:false} discounted={ID:7 Amount:850 Paid:false}/false source=[10 20 30 40] window=[20 30 777]/false`
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
