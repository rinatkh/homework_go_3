package slices

// В этом блоке нужно закрепить слайсы: append, len, cap, copy, под-слайсы,
// удаление, вставку, фильтрацию и сортировку.

func NewSlice(numbers ...int) []int {
	// TODO: создать независимый слайс из переданных чисел.
	return nil
}

func AppendValue(numbers []int, value int) []int {
	// TODO: добавить значение в конец слайса.
	return nil
}

func Sum(numbers []int) int {
	// TODO: посчитать сумму элементов.
	return 0
}

func Average(numbers []int) (float64, bool) {
	// TODO: вернуть среднее и true. Для пустого слайса вернуть 0 и false.
	return 0, false
}

func FilterEven(numbers []int) []int {
	// TODO: вернуть только чётные числа.
	return nil
}

func MapDouble(numbers []int) []int {
	// TODO: вернуть новый слайс, где каждый элемент умножен на 2.
	return nil
}

func FindIndex(numbers []int, target int) int {
	// TODO: вернуть индекс target или -1.
	return 0
}

func RemoveAt(numbers []int, index int) []int {
	// TODO: удалить элемент по индексу и вернуть новый слайс.
	return nil
}

func InsertAt(numbers []int, index int, value int) []int {
	// TODO: вставить value по index. Индекс меньше 0 считать 0, больше len — len.
	return nil
}

func CopySlice(numbers []int) []int {
	// TODO: вернуть независимую копию слайса.
	return nil
}

func ReverseInPlace(numbers []int) {
	// TODO: перевернуть слайс на месте.
}

func Unique(numbers []int) []int {
	// TODO: вернуть уникальные числа с сохранением порядка.
	return nil
}

func Window(numbers []int, start, end int) []int {
	// TODO: безопасно вернуть копию под-слайса numbers[start:end].
	return nil
}

func Chunk(numbers []int, size int) [][]int {
	// TODO: разбить слайс на части размера size.
	return nil
}

func MergeAndSort(left, right []int) []int {
	// TODO: объединить два слайса и отсортировать результат по возрастанию.
	return nil
}
