package common

import "fmt"

func Example() string {
	report, reportErr := ActiveUserReport([]User{
		{ID: 1, Name: "Maria", Active: true},
		{ID: 2, Name: "Alex"},
		{ID: 3, Name: "Ira", Active: true},
	})
	originalOrder := Order{ID: 7, Amount: 1000}
	discounted, discountErr := DiscountOrder(originalOrder, 15)
	source, window, windowErr := AppendToIsolatedWindow([]int{10, 20, 30, 40}, 1, 3, 777)

	return fmt.Sprintf(
		"report=%q/%t order=%+v discounted=%+v/%t source=%v window=%v/%t",
		report,
		reportErr != nil,
		originalOrder,
		discounted,
		discountErr != nil,
		source,
		window,
		windowErr != nil,
	)
}
