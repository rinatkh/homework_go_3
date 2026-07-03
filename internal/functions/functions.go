package functions

// В этом блоке нужно закрепить функции: параметры, return, несколько возвращаемых значений,
// variadic-параметры, функции как значения и замыкания.

func FullName(firstName, lastName string) string {
	// TODO: вернуть имя и фамилию через пробел без лишних пробелов по краям.
	return ""
}

func PriceWithDiscount(price float64, discountPercent float64) float64 {
	// TODO: применить скидку. Некорректные значения должны давать 0.
	return 0
}

func SafeDivide(a, b int) (int, bool) {
	// TODO: безопасно поделить a на b. При b == 0 вернуть 0, false.
	return 0, false
}

func MinMax(a, b int) (int, int) {
	// TODO: вернуть сначала минимум, потом максимум.
	return 0, 0
}

func NormalizeEmail(email string) string {
	// TODO: убрать пробелы по краям и привести email к нижнему регистру.
	return ""
}

func FormatUser(id int, name string, active bool) string {
	// TODO: вернуть строку вида "#7 Maria (active)" или "#7 Maria (inactive)".
	return ""
}

func ApplyOperation(a, b int, operation func(int, int) int) int {
	// TODO: применить operation к a и b. Если operation == nil, вернуть 0.
	return 0
}

func SumVariadic(numbers ...int) int {
	// TODO: сложить любое количество чисел.
	return 0
}

func BuildGreeting(language, name string) string {
	// TODO: ru -> "Привет, name", en -> "Hello, name", default -> "Hi, name".
	return ""
}

func ValidatePassword(password string) (bool, string) {
	// TODO: проверить длину >= 8 и наличие цифры. Вернуть false с причиной или true, "ok".
	return false, ""
}

func SplitFullName(fullName string) (string, string) {
	// TODO: разделить строку на имя и фамилию через strings.Fields.
	return "", ""
}

func CalcOrderTotal(price, delivery, discount float64) float64 {
	// TODO: посчитать итог заказа: цена со скидкой + доставка.
	return 0
}

func MakeCounter(start int) func() int {
	// TODO: вернуть функцию, которая при каждом вызове увеличивает счётчик на 1.
	return nil
}

func Swap(a, b string) (string, string) {
	// TODO: поменять значения местами.
	return "", ""
}

func Average(numbers ...int) (float64, bool) {
	// TODO: вернуть среднее и true. Если чисел нет, вернуть 0 и false.
	return 0, false
}
