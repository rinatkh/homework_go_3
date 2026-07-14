package functions

type User struct {
	ID     int
	Name   string
	Active bool
}

// 01. SafeDivide выполняет целочисленное деление.
// TODO: при b == 0 верните 0 и ошибку "division by zero".
func SafeDivide(a, b int) (int, error) { return 0, nil }

// 02. FindUserByID ищет пользователя по ID.
// TODO: найденный пользователь возвращается с true; отсутствие — User{} и false.
func FindUserByID(users []User, id int) (User, bool) { return User{}, false }

// 03. FindActiveUser ищет пользователя с нужным ID только среди активных.
// TODO: неактивный пользователь считается ненайденным.
func FindActiveUser(users []User, id int) (User, bool) { return User{}, false }

// 04. SumAll складывает произвольное количество чисел.
// TODO: вызов без аргументов должен вернуть 0.
func SumAll(numbers ...int) int { return 0 }

// 05. Apply вызывает переданную функцию op для a и b.
// TODO: верните результат op без изменения аргументов.
func Apply(a, b int, op func(int, int) int) int { return 0 }

// 06. ApplyIf применяет transform к value только при allowed=true.
// TODO: при false верните исходное value и не вызывайте transform.
func ApplyIf(value int, allowed bool, transform func(int) int) int { return 0 }

// 07. NewCounter создаёт замыкание-счётчик.
// TODO: первый вызов возвращённой функции должен вернуть start+1,
// каждый следующий — ещё на единицу больше.
func NewCounter(start int) func() int { return func() int { return 0 } }

// 08. NewAccumulator создаёт замыкание с накопленной суммой.
// TODO: каждый вызов добавляет аргумент к текущему состоянию и возвращает сумму.
func NewAccumulator(initial int) func(int) int { return func(int) int { return 0 } }

// 09. MakeMultiplier возвращает функцию умножения на factor.
// TODO: factor должен сохраняться внутри возвращённой функции.
func MakeMultiplier(factor int) func(int) int { return func(int) int { return 0 } }

// 10. DeferOrder возвращает строку, показывающую порядок defer.
// TODO: тело добавляет "body", затем зарегистрированы defer "first" и "second".
// Итоговая строка должна быть "body-second-first".
func DeferOrder() (result string) { return "" }

// 11. CaptureDeferArgument демонстрирует вычисление аргумента defer сразу.
// TODO: сначала value="first", затем зарегистрируйте defer с value как параметром,
// после регистрации поменяйте value на "second". Верните захваченное значение.
func CaptureDeferArgument() (result string) { return "" }

// 12. ReadDeferredVariable демонстрирует чтение переменной deferred-замыканием.
// TODO: deferred-функция без параметров должна прочитать value после того,
// как оно изменилось с "first" на "second".
func ReadDeferredVariable() (result string) { return "" }

// 13. IncrementNamedResult возвращает value, увеличенный deferred-функцией на 1.
// TODO: используйте именованное возвращаемое значение.
func IncrementNamedResult(value int) (result int) { return 0 }

// 14. RunWithCleanup выполняет action, а cleanup откладывает до выхода.
// TODO: верните две строки в порядке фактического выполнения: action, cleanup.
func RunWithCleanup(action, cleanup func() string) (events []string) { return nil }

// 15. ChooseOperation возвращает функцию для "add", "sub" или "mul".
// TODO: для неизвестного имени верните nil и false.
func ChooseOperation(name string) (func(int, int) int, bool) {
	return nil, false
}
