package arrays

// В этом блоке нужно закрепить массивы: фиксированная длина, zero value,
// индексирование, value semantics и сравнение массивов.

func ZeroArray() [3]int {
	// TODO: вернуть массив из трёх int со zero value.
	return [3]int{}
}

func Weekdays() [7]string {
	// TODO: вернуть массив дней недели: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	return [7]string{}
}

func SetAt(numbers [5]int, index int, value int) [5]int {
	// TODO: вернуть копию массива с изменённым элементом. Некорректный index игнорировать.
	return [5]int{}
}

func Sum(numbers [5]int) int {
	// TODO: посчитать сумму элементов массива.
	return 0
}

func Average(numbers [5]int) float64 {
	// TODO: посчитать среднее значение массива.
	return 0
}

func Max(numbers [5]int) int {
	// TODO: вернуть максимальный элемент массива.
	return 0
}

func Reverse(numbers [5]int) [5]int {
	// TODO: вернуть новый массив в обратном порядке.
	return [5]int{}
}

func Contains(numbers [5]int, target int) bool {
	// TODO: проверить, есть ли target в массиве.
	return false
}

func CountValue(numbers [5]int, target int) int {
	// TODO: посчитать количество вхождений target.
	return 0
}

func Equal(left, right [5]int) bool {
	// TODO: сравнить массивы.
	return false
}

func CopyAndSet(numbers [5]int, index int, value int) ([5]int, [5]int) {
	// TODO: вернуть исходный массив и изменённую копию.
	return [5]int{}, [5]int{}
}

func FirstLast(numbers [5]int) (int, int) {
	// TODO: вернуть первый и последний элемент.
	return 0, 0
}

func ToSlice(numbers [5]int) []int {
	// TODO: вернуть независимую слайс-копию массива.
	return nil
}

func MatrixDiagonalSum(matrix [3][3]int) int {
	// TODO: посчитать сумму главной диагонали.
	return 0
}

func CompareBySum(left, right [5]int) string {
	// TODO: вернуть "left", "right" или "equal" в зависимости от суммы элементов.
	return ""
}
