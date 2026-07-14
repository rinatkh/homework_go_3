package structs

type User struct {
	ID     int
	Name   string
	Active bool
}

type Product struct {
	Name  string
	Price int
}

type Order struct {
	ID     int
	Amount int
	Paid   bool
}

// 01. NewUser создаёт пользователя с ID и именем.
// TODO: Active должен получить zero value.
func NewUser(id int, name string) User { return User{} }

// 02. RenameUser возвращает копию пользователя с новым именем.
// TODO: ID и Active должны сохраниться.
func RenameUser(user User, name string) User { return User{} }

// 03. ActivateUser возвращает активную копию пользователя.
// TODO: остальные поля не изменяйте.
func ActivateUser(user User) User { return User{} }

// 04. DeactivateUser возвращает неактивную копию пользователя.
// TODO: остальные поля не изменяйте.
func DeactivateUser(user User) User { return User{} }

// 05. IsActive возвращает состояние пользователя.
// TODO: прочитайте соответствующее поле структуры.
func IsActive(user User) bool { return false }

// 06. EmptyUser возвращает zero value структуры User.
// TODO: не заполняйте поля вручную значениями, отличными от zero value.
func EmptyUser() User { return User{} }

// 07. SameUserID сравнивает пользователей только по ID.
// TODO: имя и Active не должны влиять на результат.
func SameUserID(a, b User) bool { return false }

// 08. UserLabel формирует строку "<ID>:<Name>".
// TODO: между двоеточием и именем пробел не нужен.
func UserLabel(user User) string { return "" }

// 09. NewProduct создаёт товар с названием и ценой.
// TODO: перенесите оба аргумента в соответствующие поля.
func NewProduct(name string, price int) Product { return Product{} }

// 10. ChangePrice возвращает копию товара с новой ценой.
// TODO: название товара должно сохраниться.
func ChangePrice(product Product, price int) Product { return Product{} }

// 11. ProductTotal считает стоимость count единиц товара.
// TODO: функция не должна изменять product.
func ProductTotal(product Product, count int) int { return 0 }

// 12. ApplyDiscount возвращает товар с уменьшенной ценой.
// TODO: percent находится в диапазоне от 0 до 100. Результат вычисляется
// целочисленной арифметикой, название товара сохраняется.
func ApplyDiscount(product Product, percent int) Product { return Product{} }

// 13. NewOrder создаёт заказ с ID и суммой.
// TODO: Paid должен получить zero value.
func NewOrder(id, amount int) Order { return Order{} }

// 14. MarkPaid возвращает оплаченную копию заказа.
// TODO: ID и Amount должны сохраниться.
func MarkPaid(order Order) Order { return Order{} }

// 15. OrderStatus возвращает "paid" для оплаченного заказа
// и "pending" для неоплаченного.
// TODO: выберите строку по полю Paid.
func OrderStatus(order Order) string { return "" }
