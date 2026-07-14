package structs

import "testing"

func TestNewUser(t *testing.T) {
	cases := []struct {
		name     string
		id       int
		userName string
		want     User
	}{
		{"maria", 1, "Maria", User{ID: 1, Name: "Maria"}}, {"zero id", 0, "Zero", User{Name: "Zero"}},
		{"negative id", -1, "Temp", User{ID: -1, Name: "Temp"}}, {"empty name", 2, "", User{ID: 2}},
		{"unicode", 3, "Мария", User{ID: 3, Name: "Мария"}}, {"emoji", 4, "🙂", User{ID: 4, Name: "🙂"}},
		{"space", 5, " ", User{ID: 5, Name: " "}}, {"large id", 9999, "User", User{ID: 9999, Name: "User"}},
		{"long", 6, "Long User Name", User{ID: 6, Name: "Long User Name"}}, {"another", 7, "Alex", User{ID: 7, Name: "Alex"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewUser(tc.id, tc.userName); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestRenameUser(t *testing.T) {
	cases := []struct {
		name    string
		in      User
		newName string
		want    User
	}{
		{"basic", User{ID: 1, Name: "M", Active: true}, "Maria", User{ID: 1, Name: "Maria", Active: true}},
		{"to empty", User{ID: 2, Name: "A"}, "", User{ID: 2}}, {"unicode", User{ID: 3, Name: "M"}, "Мария", User{ID: 3, Name: "Мария"}},
		{"same", User{ID: 4, Name: "Alex", Active: true}, "Alex", User{ID: 4, Name: "Alex", Active: true}},
		{"inactive", User{ID: 5, Name: "Old"}, "New", User{ID: 5, Name: "New"}}, {"zero", User{}, "Zero", User{Name: "Zero"}},
		{"emoji", User{ID: 6, Name: "x"}, "🙂", User{ID: 6, Name: "🙂"}}, {"space", User{ID: 7, Name: "x"}, " ", User{ID: 7, Name: " "}},
		{"negative id", User{ID: -1, Name: "x", Active: true}, "y", User{ID: -1, Name: "y", Active: true}},
		{"long", User{ID: 8, Name: "x"}, "Long Name", User{ID: 8, Name: "Long Name"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := RenameUser(input, tc.newName); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
			if input != tc.in {
				t.Fatal("input changed")
			}
		})
	}
}

func TestActivateUser(t *testing.T) {
	cases := []User{{}, {ID: 1}, {Name: "M"}, {Active: true}, {ID: 2, Name: "A"}, {ID: -1, Name: "X"}, {ID: 9, Name: "🙂"}, {Name: " "}, {ID: 100, Active: false}, {ID: 7, Name: "Maria", Active: true}}
	for i, in := range cases {
		t.Run(testName(i), func(t *testing.T) {
			got := ActivateUser(in)
			if !got.Active || got.ID != in.ID || got.Name != in.Name {
				t.Fatalf("got %+v from %+v", got, in)
			}
			if in != cases[i] {
				t.Fatal("input changed")
			}
		})
	}
}

func TestDeactivateUser(t *testing.T) {
	cases := []User{{}, {ID: 1, Active: true}, {Name: "M", Active: true}, {Active: false}, {ID: 2, Name: "A", Active: true}, {ID: -1, Name: "X", Active: true}, {ID: 9, Name: "🙂", Active: true}, {Name: " ", Active: true}, {ID: 100, Active: false}, {ID: 7, Name: "Maria", Active: true}}
	for i, in := range cases {
		t.Run(testName(i), func(t *testing.T) {
			got := DeactivateUser(in)
			if got.Active || got.ID != in.ID || got.Name != in.Name {
				t.Fatalf("got %+v from %+v", got, in)
			}
			if in != cases[i] {
				t.Fatal("input changed")
			}
		})
	}
}

func TestIsActive(t *testing.T) {
	cases := []struct {
		name string
		in   User
		want bool
	}{
		{"zero", User{}, false}, {"active", User{Active: true}, true}, {"inactive", User{Active: false}, false},
		{"named active", User{Name: "M", Active: true}, true}, {"named inactive", User{Name: "A"}, false},
		{"id active", User{ID: 1, Active: true}, true}, {"id inactive", User{ID: 2}, false},
		{"negative id active", User{ID: -1, Active: true}, true}, {"unicode", User{Name: "Мария", Active: true}, true},
		{"full inactive", User{ID: 9, Name: "Alex", Active: false}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsActive(tc.in); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestEmptyUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(testName(i), func(t *testing.T) {
			if got := EmptyUser(); got != (User{}) {
				t.Fatalf("got %+v", got)
			}
		})
	}
}

func TestSameUserID(t *testing.T) {
	cases := []struct {
		name string
		a, b User
		want bool
	}{
		{"same zero", User{}, User{}, true}, {"same one", User{ID: 1}, User{ID: 1}, true},
		{"different", User{ID: 1}, User{ID: 2}, false}, {"same id different name", User{ID: 3, Name: "A"}, User{ID: 3, Name: "B"}, true},
		{"same id active differs", User{ID: 4, Active: true}, User{ID: 4}, true}, {"negative same", User{ID: -1}, User{ID: -1}, true},
		{"negative different", User{ID: -1}, User{ID: 1}, false}, {"large same", User{ID: 9999}, User{ID: 9999, Name: "X"}, true},
		{"zero vs one", User{}, User{ID: 1}, false}, {"names same ids differ", User{ID: 8, Name: "M"}, User{ID: 9, Name: "M"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SameUserID(tc.a, tc.b); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestUserLabel(t *testing.T) {
	cases := []struct {
		name string
		in   User
		want string
	}{
		{"basic", User{ID: 1, Name: "Maria"}, "1:Maria"}, {"zero", User{}, "0:"}, {"negative", User{ID: -1, Name: "Temp"}, "-1:Temp"},
		{"empty name", User{ID: 2}, "2:"}, {"unicode", User{ID: 3, Name: "Мария"}, "3:Мария"}, {"emoji", User{ID: 4, Name: "🙂"}, "4:🙂"},
		{"space", User{ID: 5, Name: " "}, "5: "}, {"active ignored", User{ID: 6, Name: "A", Active: true}, "6:A"},
		{"large", User{ID: 9999, Name: "User"}, "9999:User"}, {"long", User{ID: 7, Name: "Long Name"}, "7:Long Name"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := UserLabel(tc.in); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestNewProduct(t *testing.T) {
	cases := []struct {
		name, productName string
		price             int
		want              Product
	}{
		{"book", "book", 100, Product{Name: "book", Price: 100}}, {"zero", "free", 0, Product{Name: "free"}},
		{"negative", "refund", -10, Product{Name: "refund", Price: -10}}, {"empty", "", 50, Product{Price: 50}},
		{"unicode", "курс", 2500, Product{Name: "курс", Price: 2500}}, {"emoji", "🙂", 1, Product{Name: "🙂", Price: 1}},
		{"space", " ", 5, Product{Name: " ", Price: 5}}, {"large", "laptop", 100000, Product{Name: "laptop", Price: 100000}},
		{"one", "x", 1, Product{Name: "x", Price: 1}}, {"another", "service", 999, Product{Name: "service", Price: 999}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewProduct(tc.productName, tc.price); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestChangePrice(t *testing.T) {
	cases := []struct {
		name  string
		in    Product
		price int
		want  Product
	}{
		{"increase", Product{Name: "book", Price: 100}, 200, Product{Name: "book", Price: 200}},
		{"decrease", Product{Name: "book", Price: 100}, 50, Product{Name: "book", Price: 50}},
		{"zero", Product{Name: "free", Price: 10}, 0, Product{Name: "free"}},
		{"negative", Product{Name: "refund", Price: 10}, -5, Product{Name: "refund", Price: -5}},
		{"same", Product{Name: "x", Price: 7}, 7, Product{Name: "x", Price: 7}},
		{"empty name", Product{Price: 9}, 3, Product{Price: 3}},
		{"unicode", Product{Name: "курс", Price: 1000}, 800, Product{Name: "курс", Price: 800}},
		{"large", Product{Name: "car", Price: 100}, 1000000, Product{Name: "car", Price: 1000000}},
		{"from negative", Product{Name: "x", Price: -1}, 1, Product{Name: "x", Price: 1}},
		{"space", Product{Name: " ", Price: 5}, 6, Product{Name: " ", Price: 6}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.in
			if got := ChangePrice(input, tc.price); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
			if input != tc.in {
				t.Fatal("input changed")
			}
		})
	}
}

func TestProductTotal(t *testing.T) {
	cases := []struct {
		name        string
		product     Product
		count, want int
	}{
		{"three", Product{Price: 100}, 3, 300}, {"zero count", Product{Price: 100}, 0, 0}, {"one", Product{Price: 99}, 1, 99},
		{"zero price", Product{Price: 0}, 5, 0}, {"negative count", Product{Price: 10}, -2, -20}, {"negative price", Product{Price: -10}, 2, -20},
		{"both negative", Product{Price: -10}, -2, 20}, {"large", Product{Price: 2500}, 4, 10000},
		{"named", Product{Name: "course", Price: 3500}, 2, 7000}, {"count ten", Product{Price: 7}, 10, 70},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.product
			if got := ProductTotal(input, tc.count); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
			if input != tc.product {
				t.Fatal("product changed")
			}
		})
	}
}

func TestApplyDiscount(t *testing.T) {
	cases := []struct {
		name    string
		product Product
		percent int
		want    Product
	}{
		{"ten", Product{Name: "book", Price: 1000}, 10, Product{Name: "book", Price: 900}},
		{"zero", Product{Name: "book", Price: 1000}, 0, Product{Name: "book", Price: 1000}},
		{"full", Product{Name: "book", Price: 1000}, 100, Product{Name: "book", Price: 0}},
		{"twenty five", Product{Name: "x", Price: 200}, 25, Product{Name: "x", Price: 150}},
		{"truncate", Product{Name: "x", Price: 99}, 10, Product{Name: "x", Price: 89}},
		{"one", Product{Name: "x", Price: 100}, 1, Product{Name: "x", Price: 99}},
		{"fifty", Product{Name: "x", Price: 7}, 50, Product{Name: "x", Price: 3}},
		{"zero price", Product{Name: "free"}, 80, Product{Name: "free"}},
		{"large", Product{Name: "course", Price: 2500}, 20, Product{Name: "course", Price: 2000}},
		{"unicode", Product{Name: "курс", Price: 333}, 33, Product{Name: "курс", Price: 223}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.product
			if got := ApplyDiscount(input, tc.percent); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
			if input != tc.product {
				t.Fatal("input changed")
			}
		})
	}
}

func TestNewOrder(t *testing.T) {
	cases := []struct {
		name       string
		id, amount int
		want       Order
	}{
		{"basic", 1, 500, Order{ID: 1, Amount: 500}}, {"zero", 0, 0, Order{}}, {"negative id", -1, 100, Order{ID: -1, Amount: 100}},
		{"negative amount", 2, -50, Order{ID: 2, Amount: -50}}, {"large", 999, 100000, Order{ID: 999, Amount: 100000}},
		{"one", 1, 1, Order{ID: 1, Amount: 1}}, {"zero amount", 7, 0, Order{ID: 7}}, {"negative both", -2, -3, Order{ID: -2, Amount: -3}},
		{"another", 15, 2500, Order{ID: 15, Amount: 2500}}, {"max sample", 10000, 999999, Order{ID: 10000, Amount: 999999}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := NewOrder(tc.id, tc.amount); got != tc.want {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestMarkPaid(t *testing.T) {
	cases := []Order{{}, {ID: 1}, {Amount: 100}, {Paid: true}, {ID: 2, Amount: 500}, {ID: -1, Amount: 10}, {ID: 9, Amount: -20}, {ID: 100, Paid: false}, {ID: 7, Amount: 700, Paid: true}, {ID: 8, Amount: 1}}
	for i, in := range cases {
		t.Run(testName(i), func(t *testing.T) {
			got := MarkPaid(in)
			if !got.Paid || got.ID != in.ID || got.Amount != in.Amount {
				t.Fatalf("got %+v", got)
			}
			if in != cases[i] {
				t.Fatal("input changed")
			}
		})
	}
}

func TestOrderStatus(t *testing.T) {
	cases := []struct {
		name string
		in   Order
		want string
	}{
		{"zero", Order{}, "pending"}, {"paid", Order{Paid: true}, "paid"}, {"pending", Order{Paid: false}, "pending"},
		{"id paid", Order{ID: 1, Paid: true}, "paid"}, {"id pending", Order{ID: 2}, "pending"},
		{"amount paid", Order{Amount: 100, Paid: true}, "paid"}, {"amount pending", Order{Amount: 100}, "pending"},
		{"negative paid", Order{ID: -1, Amount: -2, Paid: true}, "paid"}, {"large pending", Order{ID: 999, Amount: 100000}, "pending"},
		{"full paid", Order{ID: 7, Amount: 500, Paid: true}, "paid"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OrderStatus(tc.in); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func testName(i int) string {
	const digits = "0123456789"
	return "case_" + string(digits[i])
}

func TestExampleOutput(t *testing.T) {
	want := "user={ID:7 Name:Maria Active:false} renamed={ID:7 Name:Masha Active:false} active=true sameID=true label=7:Masha product={Name:course Price:2500} discounted={Name:course Price:2000} order={ID:15 Amount:4000 Paid:true} status=paid"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
