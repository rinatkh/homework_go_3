# Homework 3 — данные, коллекции и управляющий код

Нужно реализовать функции с `TODO` внутри `internal/*`.

У каждой задачи есть:

- сигнатура функции;
- описание результата;
- важные граничные случаи;
- примеры входных и выходных данных;
- 10 test cases в соответствующем файле `*_test.go`.

Тесты можно читать, чтобы уточнять контракт функции, но сначала попробуйте решить задачу по описанию. Не меняйте сигнатуры функций и ожидаемый формат результата.

---

## 01. Errors

Файл: `internal/errors/errors.go`

### 1. `IsNil(err error) bool`

Определите, отсутствует ли ошибка. Только настоящий `nil` считается отсутствием ошибки.

Примеры: `nil -> true`, `errors.New("boom") -> false`.

### 2. `RequireText(field, value string) error`

Проверьте обязательное текстовое поле. Для пустого значения верните `*FieldError`, в котором `Field` совпадает с аргументом, а `Reason` равен `"is required"`. Непустое значение считается корректным.

Примеры: `("email", "") -> email: is required`, `("email", "a@b") -> nil`.

### 3. `ValidatePositive(field string, value int) error`

Проверьте, что число строго больше нуля. При ошибке верните `*FieldError` для соответствующего поля с причиной `"must be positive"`.

Примеры: `("age", 20) -> nil`, `("age", 0) -> age: must be positive`.

### 4. `ParseAge(raw string) (int, error)`

Преобразуйте строку в целое число. При неуспешном преобразовании ошибка должна содержать контекст `parse age`, а исходная причина должна оставаться доступной в цепочке ошибок.

Примеры: `"26" -> 26, nil`, `"abc" -> 0, error`.

### 5. `WrapOperation(operation string, err error) error`

Добавьте к ошибке название операции. Если входная ошибка равна `nil`, результат тоже должен быть `nil`. Для ненулевой ошибки текст должен содержать операцию, а исходная ошибка не должна потеряться.

Пример: `("save user", dbErr) -> "save user: ..."`, при этом стандартная проверка причины находит `dbErr`.

### 6. `NotFound(resource string) error`

Создайте ошибку отсутствующего ресурса. В тексте должно быть название ресурса, а причиной должна оставаться `ErrNotFound`.

Пример: `"user" -> "user: not found"`.

### 7. `IsNotFound(err error) bool`

Определите, находится ли `ErrNotFound` внутри ошибки. Функция должна работать для прямой, один раз обёрнутой и несколько раз обёрнутой ошибки.

Примеры: `ErrNotFound -> true`, обычная ошибка с похожим текстом -> false`.

### 8. `SameCause(err, target error) bool`

Проверьте, относится ли `err` к причине `target`. Сравнение текста недостаточно: должна учитываться цепочка ошибок.

Примеры: обёрнутая `target -> true`, другая ошибка с таким же текстом -> false`.

### 9. `FieldFrom(err error) (*FieldError, bool)`

Попробуйте получить `*FieldError` из ошибки, включая обёрнутые ошибки. Если подходящего типа нет, верните `nil, false`.

Пример: `fmt.Errorf("validate: %w", fieldErr) -> fieldErr, true`.

### 10. `FieldName(err error) string`

Верните название поля из `*FieldError`. Ошибка может быть обёрнута несколько раз. Для `nil` и ошибок другого типа верните пустую строку.

### 11. `FirstError(errs []error) error`

Верните первую ненулевую ошибку, сохраняя исходный порядок. Для пустого списка или списка из `nil` верните `nil`.

### 12. `CountErrors(errs []error) int`

Посчитайте только ненулевые ошибки. Пустой и `nil`-слайс дают результат `0`.

### 13. `ErrorText(err error) string`

Верните текст ошибки. Для `nil` верните пустую строку.

### 14. `ValidateUser(name, email string, age int) error`

Проверьте пользователя в порядке `name`, `email`, `age`. Имя и email обязательны, возраст должен быть положительным. Верните только первую найденную ошибку как `*FieldError`.

Примеры: `("Maria", "m@x", 20) -> nil`; `("", "", 0)` сообщает сначала об имени.

### 15. `Classify(err error) string`

Верните категорию ошибки:

- `nil` — `"none"`;
- `ErrNotFound` в цепочке — `"not_found"`;
- `*FieldError` — `"field:<имя поля>"`;
- другая ошибка — `"other"`.

Проверка раздела:

```bash
make test-errors
make run-01_errors
```

---

## 02. Arrays

Файл: `internal/arrays/arrays.go`

### 1. `FirstOfThree(items [3]int) int`
Верните первый элемент массива.

### 2. `LastOfFour(items [4]string) string`
Верните последний элемент массива, не привязываясь к конкретным значениям.

### 3. `SetMiddle(items [3]int, value int) [3]int`
Верните массив с заменённым средним элементом. Остальные позиции должны сохраниться.

Пример: `[1 2 3], 9 -> [1 9 3]`.

### 4. `SwapEdges(items [4]int) [4]int`
Поменяйте местами первый и последний элементы. Два внутренних элемента остаются на прежних позициях.

### 5. `SumThree(items [3]int) int`
Верните сумму всех трёх элементов.

### 6. `AverageFour(items [4]int) int`
Верните целочисленное среднее четырёх чисел. Дробная часть отбрасывается по правилам Go.

### 7. `ReverseFive(items [5]int) [5]int`
Верните элементы в обратном порядке.

### 8. `EqualPairs(a, b [2]string) bool`
Сравните массивы целиком. Порядок и регистр строк имеют значение.

### 9. `CopyAndSet(items [3]int, index, value int) ([3]int, [3]int)`
Верните два массива: исходную копию и копию, где изменён только элемент `index`. В тестах индекс всегда корректен.

### 10. `CountTrue(items [5]bool) int`
Посчитайте количество значений `true`.

### 11. `ContainsFour(items [4]int, target int) bool`
Проверьте наличие целевого числа в массиве.

### 12. `MaxFour(items [4]int) int`
Найдите максимальное значение. Учтите массивы, состоящие только из отрицательных чисел.

### 13. `MainDiagonalSum(matrix [3][3]int) int`
Верните сумму элементов главной диагонали квадратной матрицы 3×3.

### 14. `CornersSum(matrix [3][3]int) int`
Верните сумму четырёх углов матрицы. Центральный элемент и середины сторон не учитываются.

### 15. `RotateRight(items [4]int) [4]int`
Сдвиньте массив вправо на одну позицию: последний элемент переходит в начало.

Пример: `[1 2 3 4] -> [4 1 2 3]`.

Проверка:

```bash
make test-arrays
make run-02_arrays
```

---

## 03. Structs

Файл: `internal/structs/structs.go`

В этом разделе структуры используются только как данные: без методов и интерфейсов.

### 1. `NewUser(id int, name string) User`
Создайте пользователя с ID и именем. `Active` получает zero value.

### 2. `RenameUser(user User, name string) User`
Верните копию пользователя с новым именем. ID и активность сохраняются, исходный аргумент не изменяется.

### 3. `ActivateUser(user User) User`
Верните активную копию пользователя, сохранив остальные поля.

### 4. `DeactivateUser(user User) User`
Верните неактивную копию пользователя, сохранив остальные поля.

### 5. `IsActive(user User) bool`
Верните текущее состояние пользователя.

### 6. `EmptyUser() User`
Верните zero value структуры `User`.

### 7. `SameUserID(a, b User) bool`
Сравните пользователей только по ID. Имя и активность не влияют на ответ.

### 8. `UserLabel(user User) string`
Сформируйте строку `<ID>:<Name>` без пробела после двоеточия.

### 9. `NewProduct(name string, price int) Product`
Создайте товар с названием и ценой.

### 10. `ChangePrice(product Product, price int) Product`
Верните копию товара с новой ценой. Название и исходный аргумент сохраняются.

### 11. `ProductTotal(product Product, count int) int`
Посчитайте итоговую стоимость указанного количества товара.

### 12. `ApplyDiscount(product Product, percent int) Product`
Верните копию товара со скидкой. Процент находится в диапазоне 0–100. Вычисление выполняется целыми числами.

Пример: цена `1000`, скидка `10` -> цена `900`.

### 13. `NewOrder(id, amount int) Order`
Создайте заказ с ID и суммой. `Paid` получает zero value.

### 14. `MarkPaid(order Order) Order`
Верните оплаченную копию заказа, сохранив ID и сумму.

### 15. `OrderStatus(order Order) string`
Верните `"paid"` для оплаченного заказа и `"pending"` для неоплаченного.

Проверка:

```bash
make test-structs
make run-03_structs
```

---

## 04. New / Make

Файл: `internal/newmake/newmake.go`

### 1. `NewInt(value int) *int`
Создайте ненулевой указатель на `int` и запишите в него значение.

### 2. `NewString(value string) *string`
Создайте ненулевой указатель на `string` с заданным значением.

### 3. `NewZeroUser() *User`
Верните ненулевой указатель на zero value `User`.

### 4. `NewNamedUser(name string) *User`
Верните указатель на пользователя с заполненным именем. `Scores` остаётся nil.

### 5. `MakeInts(length int) []int`
Создайте ненулевой слайс указанной длины. Все элементы получают zero value.

### 6. `MakeIntsWithCapacity(length, capacity int) []int`
Создайте слайс с указанными `len` и `cap`. В тестах `capacity >= length`.

### 7. `MakeAndAppend(capacity int, values ...int) []int`
Создайте слайс с начальной вместимостью и добавьте все значения, сохранив порядок.

### 8. `MakeEmptyStrings() []string`
Создайте пустой, но не nil слайс строк.

### 9. `SliceState(items []int) (length, capacity int, isNil bool)`
Верните состояние входного слайса: длину, вместимость и результат сравнения с nil.

### 10. `EmptySliceState() (length, capacity int, isNil bool)`
Создайте пустой слайс через `make` и верните его состояние.

Ожидаемая логика: длина 0, вместимость 0, `isNil=false`.

### 11. `MakeUserScores(name string, count int) *User`
Создайте пользователя и ненулевой слайс `Scores` заданной длины.

### 12. `NewSlicePointer() *[]int`
Верните результат создания zero value для типа `[]int` через `new`. Сам указатель должен быть ненулевым, а слайс по этому указателю — nil.

### 13. `MakeSlicePointer(capacity int) *[]int`
Создайте пустой ненулевой слайс с заданной вместимостью и верните указатель на него.

### 14. `MakeStringIntMap() map[string]int`
Создайте пустую рабочую map, в которую можно записывать значения без panic.

### 15. `MakeBoolChannel(buffer int) chan bool`
Создайте канал `bool` с заданным размером буфера. При `buffer=0` канал должен быть небуферизованным.

Проверка:

```bash
make test-newmake
make run-04_new_make
```

---

## 05. Slices

Файл: `internal/slices/slices.go`

Это основной практический блок урока. Обратите внимание на общий базовый массив, формулы `len/cap`, поведение `append` и полный срез.

### 1. `First(items []int) (int, bool)`
Верните первый элемент и `true`. Для nil и пустого слайса верните `0, false`.

### 2. `Last(items []string) (string, bool)`
Верните последний элемент и `true`. Для nil и пустого слайса верните `"", false`.

### 3. `SliceInfo(items []int, low, high int) (part []int, length, capacity int)`
Создайте обычный срез `items[low:high]` и верните сам срез, его длину и вместимость. Границы в тестах корректны.

Пример: для `[10 20 30 40]` и границ `1:3` результат `[20 30]`, `len=2`, `cap=3`.

### 4. `FullSliceInfo(items []int, low, high, max int) (part []int, length, capacity int)`
Создайте полный срез `items[low:high:max]` и верните его состояние.

Пример: `[10 20 30 40]`, `1:3:3` -> `[20 30]`, `len=2`, `cap=2`.

### 5. `ChangeFirst(part []int, value int)`
Измените первый элемент переданного слайса. Пустой слайс должен остаться без изменений и не приводить к panic.

### 6. `MutateWindowInFunction(items []int, low, high, value int) []int`
Создайте окно `items[low:high]`, передайте его в `ChangeFirst` и верните `items`. Изменение через окно должно быть видно в исходном слайсе. Пустое окно ничего не меняет.

Пример: `[10 20 30 40]`, окно `1:3`, значение `999` -> `[10 999 30 40]`.

### 7. `AppendWindowInFunction(items []int, low, high, value int) (source, part []int)`
Создайте обычное окно, добавьте в него значение и верните исходный слайс вместе с новым окном. Нужно получить естественное поведение `append`:

- если у окна хватает capacity, следующий элемент исходного массива может быть перезаписан;
- если capacity закончилась, окно переходит на новый массив, а исходный слайс не получает новый элемент.

Пример с доступной capacity: `[10 20 30 40]`, окно `1:3`, `777` -> source `[10 20 30 777]`, part `[20 30 777]`.

### 8. `AppendLimitedWindowInFunction(items []int, low, high, value int) (source, part []int)`
Создайте окно, у которого capacity заканчивается на `high`, затем добавьте значение. Новый элемент не должен перезаписывать элемент исходного слайса за границей окна.

Пример: source остаётся `[10 20 30 40]`, part становится `[20 30 777]`.

### 9. `Clone(items []int) []int`
Верните независимую копию значений. Изменение результата не должно менять исходный слайс. Для nil верните nil; для пустого ненулевого слайса — пустой ненулевой слайс.

### 10. `ChangeClone(items []int, index, value int) (source, clone []int)`
Создайте независимую копию, измените в ней один элемент и верните исходный слайс вместе с копией. В тестах индекс корректен.

### 11. `AppendOne(items []int, value int) []int`
Добавьте один элемент и верните результат, сохранив порядок.

### 12. `AppendMany(items []int, values ...int) []int`
Добавьте любое количество элементов. Вызов без добавляемых значений также должен быть корректным.

### 13. `CanAppendWithoutGrow(items []int, extra int) bool`
Определите, можно ли добавить `extra` элементов, не превышая текущую capacity. Отрицательное `extra` считается некорректным и даёт `false`.

Примеры: `len=2, cap=5, extra=3 -> true`; `extra=4 -> false`.

### 14. `SliceKind(items []int) string`
Верните:

- `"nil"` для nil-слайса;
- `"empty"` для ненулевого слайса длины 0;
- `"filled"` для непустого слайса.

### 15. `AppendIndependent(items []int, values ...int) []int`
Верните независимый слайс, содержащий сначала `items`, затем `values`. Исходные значения и его базовый массив не должны измениться даже при свободной capacity.

Проверка:

```bash
make test-slices
make run-05_slices
```

---

## 06. Loops

Файл: `internal/loops/loops.go`

### 1. `SumTo(n int) int`
Посчитайте сумму от 1 до `n` включительно. Для `n <= 0` верните 0.

### 2. `SumBetween(start, end int) int`
Посчитайте сумму всех целых чисел в диапазоне включительно. Если `start > end`, верните 0.

### 3. `CountDown(n int) []int`
Верните числа от `n` до 1. Для `n <= 0` верните пустой ненулевой слайс.

### 4. `Factorial(n int) int`
Вычислите факториал циклом. `0!` и `1!` равны 1; для отрицательного числа верните 0.

### 5. `CountEven(items []int) int`
Посчитайте чётные значения. Ноль считается чётным.

### 6. `FirstNegative(items []int) (int, bool)`
Найдите первое отрицательное число. Если его нет, верните `0, false`.

### 7. `SumWithoutZeros(items []int) int`
Сложите все значения, пропуская нули.

### 8. `SumUntilLimit(items []int, limit int) int`
Добавляйте элементы по порядку. Если следующий элемент сделал бы сумму больше `limit`, остановитесь и верните уже накопленное значение. Для отрицательного лимита верните 0.

Пример: `[3 4 5]`, limit `7` -> `7`.

### 9. `DoubleInPlace(items []int)`
Умножьте каждый элемент на 2 непосредственно в исходном слайсе.

### 10. `ReplaceNegativeInPlace(items []int, replacement int)`
Замените все отрицательные элементы указанным значением. Ноль и положительные числа сохраняются.

### 11. `CountActive(users []User) int`
Посчитайте активных пользователей.

### 12. `ActiveNames(users []User) []string`
Верните имена активных пользователей в исходном порядке. Для пустого входа верните пустой ненулевой слайс.

### 13. `RuneCount(text string) int`
Посчитайте Unicode-символы через обход строки. Кириллица и emoji должны считаться символами, а не отдельными байтами.

### 14. `RuneByteIndexes(text string) []int`
Верните байтовые индексы начала каждой rune в строке.

Пример: `"Яa🙂" -> [0 2 3]`.

### 15. `RepeatEach(items []int, times int) []int`
Повторите каждый элемент `times` раз подряд. Для `times <= 0` верните пустой ненулевой слайс.

Пример: `[1 2], 2 -> [1 1 2 2]`.

Проверка:

```bash
make test-loops
make run-06_loops
```

---

## 07. Functions / Defer

Файл: `internal/functions/functions.go`

### 1. `SafeDivide(a, b int) (int, error)`
Выполните целочисленное деление. При делении на ноль верните `0` и ошибку с текстом `division by zero`.

### 2. `FindUserByID(users []User, id int) (User, bool)`
Найдите первого пользователя с указанным ID. Активность не влияет на поиск.

### 3. `FindActiveUser(users []User, id int) (User, bool)`
Найдите первого активного пользователя с указанным ID. Неактивный пользователь считается ненайденным.

### 4. `SumAll(numbers ...int) int`
Сложите произвольное количество чисел. Вызов без аргументов возвращает 0.

### 5. `Apply(a, b int, op func(int, int) int) int`
Вызовите переданную функцию для `a` и `b` и верните её результат.

### 6. `ApplyIf(value int, allowed bool, transform func(int) int) int`
Примените преобразование только при `allowed=true`. При `false` верните исходное значение, не вызывая `transform`.

### 7. `NewCounter(start int) func() int`
Создайте замыкание-счётчик. Первый вызов возвращает `start+1`, каждый следующий увеличивает состояние ещё на единицу. Разные счётчики не должны делить состояние.

### 8. `NewAccumulator(initial int) func(int) int`
Создайте замыкание, которое добавляет каждый новый аргумент к накопленной сумме и возвращает текущее значение.

### 9. `MakeMultiplier(factor int) func(int) int`
Верните функцию, умножающую аргумент на сохранённый `factor`.

### 10. `DeferOrder() string`
Получите строку `body-second-first`, используя тело функции и два отложенных действия `first` и `second`. Задача проверяет обратный порядок выполнения defer.

### 11. `CaptureDeferArgument() string`
Покажите, что параметр отложенного вызова вычисляется в момент регистрации. Переменная сначала равна `first`, после регистрации меняется на `second`, но функция должна вернуть `first`.

### 12. `ReadDeferredVariable() string`
Покажите другой вариант: deferred-замыкание без параметра читает переменную при фактическом выполнении. Функция должна вернуть `second`.

### 13. `IncrementNamedResult(value int) int`
Верните значение, которое увеличивается на единицу отложенным действием перед фактическим выходом из функции.

### 14. `RunWithCleanup(action, cleanup func() string) []string`
Выполните `action`, а `cleanup` отложите до выхода. Верните строки в реальном порядке выполнения: сначала результат action, затем cleanup.

### 15. `ChooseOperation(name string) (func(int, int) int, bool)`
Верните функцию для операций `add`, `sub` или `mul`. Для неизвестного имени верните `nil, false`.

Проверка:

```bash
make test-functions
make run-07_functions
```

---

## 08. Panic / Recover

Файл: `internal/panics/panics.go`

`panic` используется здесь только как учебный механизм. В обычной бизнес-валидации следует возвращать `error`.

### 1. `SafeRun(fn func()) bool`
Запустите функцию и верните `true`, если она завершилась panic. Panic не должна выйти наружу.

### 2. `RecoverMessage(fn func()) string`
Верните текст значения, переданного в panic. Если panic не было, верните пустую строку.

### 3. `PanicIfEmpty(text string)`
Для пустой строки вызовите panic со значением `empty text`. Любая непустая строка должна завершаться нормально.

### 4. `MustPositive(n int) int`
Верните положительное число без изменений. Для нуля и отрицательных значений вызовите panic `not positive`.

### 5. `ParseOrPanic(raw string) int`
Преобразуйте строку в целое число. При ошибке преобразования вызовите panic с исходной ошибкой.

### 6. `TryParse(raw string) (int, bool)`
Используйте поведение `ParseOrPanic`, но преобразуйте panic в `0, false`. Корректное число возвращается с `true`.

### 7. `SafeIndex(items []int, index int) (int, bool)`
Верните значение по индексу. Выход за границы не должен приводить к panic за пределами функции: верните `0, false`.

### 8. `MustGet(items []string, index int) string`
Верните строку по индексу и не перехватывайте естественную panic при неверном индексе.

### 9. `SafeMustGet(items []string, index int) (string, bool)`
Вызовите `MustGet`, но преобразуйте panic в `"", false`.

### 10. `PanicToError(fn func()) error`
Преобразуйте panic в ошибку формата `panic: <значение>`. Нормальное выполнение даёт `nil`.

### 11. `DeferBeforePanic() []string`
Покажите, что defer выполняется при panic. Итоговая последовательность должна быть `[body defer]`, а panic не должна выйти наружу.

### 12. `RecoverOutsideDefer() bool`
Вызовите `recover` в обычном коде и верните, получил ли он ненулевое значение. Ожидаемый результат — `false`.

### 13. `RunSequence(functions []func()) int`
Выполните функции по порядку без recover. Если все завершились нормально, верните их количество. Если одна паникует, выполнение останавливается естественно, следующие функции не запускаются.

### 14. `RunSequenceSafe(functions []func()) (completed, panicked int)`
Выполните все функции. Отдельная panic не должна останавливать оставшуюся последовательность. Верните количество нормальных завершений и panic.

### 15. `MustNonNil(value *int) int`
Верните число по указателю. Для nil вызовите panic `nil pointer`.

Проверка:

```bash
make test-panics
make run-08_panics
```

---

## 09. Общие задачи

Файл: `internal/common/common.go`

### 1. `ActiveUserReport(users []User) (string, error)`

Соберите имена активных пользователей в исходном порядке и сформируйте строку:

```text
active=<количество> names=<имя1>,<имя2>
```

Если активных пользователей нет, верните пустую строку и `ErrNoActiveUsers`.

Пример: Maria active, Alex inactive, Ira active -> `active=2 names=Maria,Ira`.

### 2. `DiscountOrder(order Order, percent int) (Order, error)`

Верните копию заказа с применённой скидкой. Исходный заказ не меняется.

Правила:

- отрицательная сумма — `*FieldError` для поля `amount`;
- процент вне диапазона 0–100 — `*FieldError` для поля `percent`;
- `ID` и `Paid` сохраняются;
- вычисление выполняется целыми числами.

### 3. `AppendToIsolatedWindow(items []int, low, high int, values ...int) (source, window []int, err error)`

Проверьте границы `0 <= low <= high <= len(items)`. При ошибке верните ошибку и nil-слайсы.

При корректных границах создайте окно, добавьте значения и верните исходный слайс вместе с новым окном. Добавление не должно перезаписывать данные исходного слайса за границей `high`.

Пример: `[10 20 30 40]`, окно `1:3`, значение `777` -> source `[10 20 30 40]`, window `[20 30 777]`.

Проверка:

```bash
make test-common
make run-09_common
```

---

## Интеграционные тесты

`Example()` каждого раздела вызывает не одну заглушку, а большую часть функций раздела. Интеграционные тесты запускают настоящие программы через `go run` и сравнивают полный вывод.

После unit-тестов обязательно выполните:

```bash
make test-integration
make run-all
```

## Финальная проверка

```bash
make fmt
make test-unit
make test-integration
make test-race
make coverage-check
make ci
```
