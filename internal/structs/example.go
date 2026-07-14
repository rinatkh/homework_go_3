package structs

import "fmt"

func Example() string {
	user := NewUser(7, "Maria")
	renamed := RenameUser(user, "Masha")
	active := ActivateUser(renamed)
	product := NewProduct("course", 2500)
	discounted := ApplyDiscount(product, 20)
	order := MarkPaid(NewOrder(15, ProductTotal(discounted, 2)))

	return fmt.Sprintf(
		"user=%+v renamed=%+v active=%t sameID=%t label=%s product=%+v discounted=%+v order=%+v status=%s",
		user,
		renamed,
		IsActive(active),
		SameUserID(user, active),
		UserLabel(active),
		product,
		discounted,
		order,
		OrderStatus(order),
	)
}
