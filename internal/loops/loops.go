package loops

// В этом блоке нужно закрепить циклы: for, range, break, continue и вложенные циклы.
// Внутри функций оставлены заготовки. Ориентируйтесь на комментарии и unit-тесты.

func SumTo(n int) int {
	// TODO: вернуть сумму чисел от 1 до n. Если n <= 0, вернуть 0.
	return 0
}

func Factorial(n int) int {
	// TODO: вернуть n!. Для отрицательных n вернуть 0, для 0 вернуть 1.
	return 0
}

func CountEven(numbers []int) int {
	// TODO: посчитать количество чётных чисел.
	return 0
}

func FindFirstNegative(numbers []int) (int, bool) {
	// TODO: вернуть индекс первого отрицательного числа и true. Если нет — -1 и false.
	return 0, false
}

func SkipMultiplesOfThree(n int) []int {
	// TODO: вернуть числа от 1 до n, пропуская кратные 3 через continue.
	return nil
}

func MultiplicationRow(n, limit int) []int {
	// TODO: вернуть n*1, n*2, ..., n*limit. Если limit <= 0, вернуть пустой слайс.
	return nil
}

func ReverseStringByRunes(s string) string {
	// TODO: перевернуть строку по рунам, чтобы корректно работал Unicode.
	return ""
}

func CountRunes(s string) map[rune]int {
	// TODO: посчитать количество каждой руны в строке.
	return nil
}

func FizzBuzz(n int) []string {
	// TODO: вернуть FizzBuzz для чисел от 1 до n.
	return nil
}

func TriangleRows(height int) []string {
	// TODO: построить строки "*", "**", "***" ... до height.
	return nil
}

func SumUntilLimit(numbers []int, limit int) int {
	// TODO: складывать числа, пока следующая прибавка не превысит limit. Используйте break.
	return 0
}

func Flatten(matrix [][]int) []int {
	// TODO: развернуть матрицу в один слайс через вложенные циклы.
	return nil
}

func MaxInSlice(numbers []int) (int, bool) {
	// TODO: вернуть максимум и true. Для пустого слайса вернуть 0 и false.
	return 0, false
}

func UniquePreserveOrder(numbers []int) []int {
	// TODO: вернуть уникальные числа с сохранением порядка первого появления.
	return nil
}

func RepeatString(word string, times int) string {
	// TODO: повторить строку times раз. Если times <= 0, вернуть пустую строку.
	return ""
}
