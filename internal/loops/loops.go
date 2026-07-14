package loops

type User struct {
	Name   string
	Active bool
}

// 01. SumTo считает сумму целых чисел от 1 до n включительно.
// TODO: для n <= 0 верните 0.
func SumTo(n int) int { return 0 }

// 02. SumBetween считает сумму от start до end включительно.
// TODO: если start > end, верните 0.
func SumBetween(start, end int) int { return 0 }

// 03. CountDown возвращает числа от n до 1.
// TODO: для n <= 0 верните пустой ненулевой слайс.
func CountDown(n int) []int { return nil }

// 04. Factorial вычисляет n! циклом.
// TODO: 0! и 1! равны 1; для отрицательного n верните 0.
func Factorial(n int) int { return 0 }

// 05. CountEven считает чётные числа в слайсе.
// TODO: ноль считается чётным.
func CountEven(items []int) int { return 0 }

// 06. FirstNegative возвращает первое отрицательное значение и true.
// TODO: если отрицательных чисел нет, верните zero value и false.
func FirstNegative(items []int) (int, bool) { return 0, false }

// 07. SumWithoutZeros суммирует элементы, пропуская нули.
// TODO: используйте управляющую конструкцию, которая переходит
// к следующей итерации.
func SumWithoutZeros(items []int) int { return 0 }

// 08. SumUntilLimit добавляет элементы по порядку, пока следующий элемент
// не сделал бы сумму больше limit.
// TODO: в этот момент остановите цикл. Для limit < 0 верните 0.
func SumUntilLimit(items []int, limit int) int { return 0 }

// 09. DoubleInPlace умножает каждый элемент слайса на 2.
// TODO: измените исходный слайс, а не копию value из range.
func DoubleInPlace(items []int) {}

// 10. ReplaceNegativeInPlace заменяет отрицательные элементы на replacement.
// TODO: ноль и положительные значения должны сохраниться.
func ReplaceNegativeInPlace(items []int, replacement int) {}

// 11. CountActive считает активных пользователей.
// TODO: порядок пользователей не влияет на количество.
func CountActive(users []User) int { return 0 }

// 12. ActiveNames возвращает имена только активных пользователей.
// TODO: сохраните исходный порядок.
func ActiveNames(users []User) []string { return nil }

// 13. RuneCount считает Unicode-символы через range по строке.
// TODO: результат для кириллицы и emoji не должен равняться числу байт.
func RuneCount(text string) int { return 0 }

// 14. RuneByteIndexes возвращает байтовые индексы всех rune в строке.
// TODO: используйте индекс, который отдаёт range по string.
func RuneByteIndexes(text string) []int { return nil }

// 15. RepeatEach повторяет каждый элемент times раз подряд.
// TODO: для times <= 0 верните пустой ненулевой слайс.
func RepeatEach(items []int, times int) []int { return nil }
