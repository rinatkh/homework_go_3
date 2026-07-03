# Задание 3: циклы, функции, массивы, слайсы

Нужно реализовать функции с `TODO` внутри `internal/*`.

## 01. Loops

Файл: `internal/loops/loops.go`

1. `SumTo` — сумма чисел от 1 до n.
2. `Factorial` — факториал.
3. `CountEven` — количество чётных чисел.
4. `FindFirstNegative` — индекс первого отрицательного числа.
5. `SkipMultiplesOfThree` — числа от 1 до n без кратных 3.
6. `MultiplicationRow` — строка таблицы умножения.
7. `ReverseStringByRunes` — разворот строки по рунам.
8. `CountRunes` — частотный словарь рун.
9. `FizzBuzz` — классическая задача на циклы и условия.
10. `TriangleRows` — текстовый треугольник.
11. `SumUntilLimit` — суммирование до лимита через `break`.
12. `Flatten` — разворот матрицы.
13. `MaxInSlice` — максимум в слайсе.
14. `UniquePreserveOrder` — уникальные значения с сохранением порядка.
15. `RepeatString` — повтор строки через цикл.

Проверка:

```bash
make test-loops
make run-loops
```

## 02. Functions

Файл: `internal/functions/functions.go`

1. `FullName`
2. `PriceWithDiscount`
3. `SafeDivide`
4. `MinMax`
5. `NormalizeEmail`
6. `FormatUser`
7. `ApplyOperation`
8. `SumVariadic`
9. `BuildGreeting`
10. `ValidatePassword`
11. `SplitFullName`
12. `CalcOrderTotal`
13. `MakeCounter`
14. `Swap`
15. `Average`

Проверка:

```bash
make test-functions
make run-functions
```

## 03. Arrays

Файл: `internal/arrays/arrays.go`

1. `ZeroArray`
2. `Weekdays`
3. `SetAt`
4. `Sum`
5. `Average`
6. `Max`
7. `Reverse`
8. `Contains`
9. `CountValue`
10. `Equal`
11. `CopyAndSet`
12. `FirstLast`
13. `ToSlice`
14. `MatrixDiagonalSum`
15. `CompareBySum`

Проверка:

```bash
make test-arrays
make run-arrays
```

## 04. Slices

Файл: `internal/slices/slices.go`

1. `NewSlice`
2. `AppendValue`
3. `Sum`
4. `Average`
5. `FilterEven`
6. `MapDouble`
7. `FindIndex`
8. `RemoveAt`
9. `InsertAt`
10. `CopySlice`
11. `ReverseInPlace`
12. `Unique`
13. `Window`
14. `Chunk`
15. `MergeAndSort`

Проверка:

```bash
make test-slices
make run-slices
```

## Финальная проверка

```bash
make test-unit
make test-integration
make ci
```
