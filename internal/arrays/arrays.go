package arrays

// Блок закрепляет фиксированную длину массива, индексы, копирование,
// сравнение и работу с небольшими матрицами.

// 01. FirstOfThree возвращает первый элемент массива из трёх чисел.
// TODO: получите значение по правильному индексу.
func FirstOfThree(items [3]int) int { return 0 }

// 02. LastOfFour возвращает последний элемент массива из четырёх строк.
// TODO: решение не должно зависеть от конкретных значений массива.
func LastOfFour(items [4]string) string { return "" }

// 03. SetMiddle возвращает копию массива с заменённым средним элементом.
// TODO: исходный аргумент снаружи функции не должен измениться.
func SetMiddle(items [3]int, value int) [3]int { return [3]int{} }

// 04. SwapEdges меняет местами первый и последний элементы копии массива.
// TODO: два внутренних элемента должны сохранить свои позиции.
func SwapEdges(items [4]int) [4]int { return [4]int{} }

// 05. SumThree возвращает сумму трёх элементов.
// TODO: учтите все позиции массива.
func SumThree(items [3]int) int { return 0 }

// 06. AverageFour возвращает целочисленное среднее четырёх чисел.
// TODO: используйте обычные правила целочисленного деления Go.
func AverageFour(items [4]int) int { return 0 }

// 07. ReverseFive возвращает массив в обратном порядке.
// TODO: первый элемент должен стать последним, второй — предпоследним и т.д.
func ReverseFive(items [5]int) [5]int { return [5]int{} }

// 08. EqualPairs сравнивает два массива строк целиком.
// TODO: порядок элементов имеет значение.
func EqualPairs(a, b [2]string) bool { return false }

// 09. CopyAndSet возвращает исходную копию и изменённую копию массива.
// TODO: измените только элемент с индексом index во втором результате.
// В тестах index всегда находится в допустимом диапазоне.
func CopyAndSet(items [3]int, index, value int) ([3]int, [3]int) {
	return [3]int{}, [3]int{}
}

// 10. CountTrue считает true в массиве из пяти bool.
// TODO: false не увеличивает счётчик.
func CountTrue(items [5]bool) int { return 0 }

// 11. ContainsFour проверяет наличие target в массиве.
// TODO: верните true при первом совпадении, иначе false.
func ContainsFour(items [4]int, target int) bool { return false }

// 12. MaxFour возвращает максимальное значение массива.
// TODO: корректно обработайте массив, состоящий только из отрицательных чисел.
func MaxFour(items [4]int) int { return 0 }

// 13. MainDiagonalSum возвращает сумму главной диагонали матрицы 3x3.
// TODO: используйте элементы [0][0], [1][1] и [2][2].
func MainDiagonalSum(matrix [3][3]int) int { return 0 }

// 14. CornersSum возвращает сумму четырёх углов матрицы 3x3.
// TODO: центральный элемент и середины сторон не учитываются.
func CornersSum(matrix [3][3]int) int { return 0 }

// 15. RotateRight сдвигает массив вправо на одну позицию.
// TODO: последний элемент должен перейти в начало.
func RotateRight(items [4]int) [4]int { return [4]int{} }
