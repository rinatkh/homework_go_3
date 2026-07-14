package slices

// 01. First возвращает первый элемент и true.
// TODO: для nil и пустого слайса верните zero value и false.
func First(items []int) (int, bool) { return 0, false }

// 02. Last возвращает последнюю строку и true.
// TODO: для nil и пустого слайса верните пустую строку и false.
func Last(items []string) (string, bool) { return "", false }

// 03. SliceInfo создаёт обычный срез items[low:high].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func SliceInfo(items []int, low, high int) (part []int, length, capacity int) {
	return nil, 0, 0
}

// 04. FullSliceInfo создаёт полный срез items[low:high:max].
// TODO: верните сам срез, его len и cap. В тестах границы корректны.
func FullSliceInfo(items []int, low, high, max int) (part []int, length, capacity int) {
	return nil, 0, 0
}

// 05. ChangeFirst меняет первый элемент переданного слайса.
// TODO: пустой слайс оставьте без изменений.
func ChangeFirst(part []int, value int) {}

// 06. MutateWindowInFunction создаёт срез items[low:high], передаёт его
// в ChangeFirst и возвращает исходный items.
// TODO: изменение должно быть видно в исходном слайсе. Пустое окно не меняйте.
func MutateWindowInFunction(items []int, low, high, value int) []int { return nil }

// 07. AppendWindowInFunction создаёт обычный срез, добавляет value
// и возвращает исходный слайс вместе с получившимся окном.
// TODO: поведение должно зависеть от доступной capacity: append может
// изменить исходный массив или перейти на новый.
func AppendWindowInFunction(items []int, low, high, value int) (source, part []int) {
	return nil, nil
}

// 08. AppendLimitedWindowInFunction создаёт окно с capacity, ограниченной high,
// затем добавляет value и возвращает исходный слайс и новое окно.
// TODO: добавление не должно перезаписывать элемент исходного слайса за high.
func AppendLimitedWindowInFunction(items []int, low, high, value int) (source, part []int) {
	return nil, nil
}

// 09. Clone возвращает независимую копию значений.
// TODO: изменение результата не должно менять items. Для nil верните nil.
func Clone(items []int) []int { return nil }

// 10. ChangeClone создаёт независимую копию, меняет в ней элемент index
// и возвращает исходный слайс и копию.
// TODO: в тестах index корректен.
func ChangeClone(items []int, index, value int) (source, clone []int) {
	return nil, nil
}

// 11. AppendOne добавляет один элемент и возвращает результат append.
// TODO: сохраните исходный порядок элементов.
func AppendOne(items []int, value int) []int { return nil }

// 12. AppendMany добавляет все values в исходном порядке.
// TODO: корректно обработайте пустой список добавляемых значений.
func AppendMany(items []int, values ...int) []int { return nil }

// 13. CanAppendWithoutGrow сообщает, хватает ли текущей capacity для extra
// новых элементов.
// TODO: отрицательное extra считается некорректным и даёт false.
func CanAppendWithoutGrow(items []int, extra int) bool { return false }

// 14. SliceKind классифицирует слайс.
// TODO: nil -> "nil", ненулевой слайс с len=0 -> "empty",
// остальные значения -> "filled".
func SliceKind(items []int) string { return "" }

// 15. AppendIndependent возвращает независимый результат items+values.
// TODO: исходный слайс и его базовый массив не должны измениться,
// даже если в исходной capacity есть свободное место.
func AppendIndependent(items []int, values ...int) []int { return nil }
