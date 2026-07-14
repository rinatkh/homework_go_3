# Homework Go 3 — данные, коллекции и управляющий код

Домашняя работа закрепляет только темы третьего занятия:

- `error`, `errors.Is`, `errors.As`, оборачивание ошибок;
- массивы и копирование массивов;
- структуры без методов;
- `new` и `make`;
- слайсы, `len`, `cap`, `append`, `copy`, общий базовый массив;
- `for`, `range`, `break`, `continue`;
- функции, несколько возвращаемых значений, variadic, функции как значения, closure;
- `defer`, порядок LIFO и момент вычисления аргументов;
- `panic` и `recover`.

В каждом основном разделе находится **15 задач**. Для каждой функции подготовлено **10 test cases** с обычными и граничными входными данными. В конце находятся 3 объединяющие задачи.

## Структура

- `docs/task.md` — полное описание всех задач;
- `internal/<topic>/<topic>.go` — функции с `TODO`;
- `internal/<topic>/<topic>_test.go` — unit-тесты;
- `internal/<topic>/example.go` — итоговый пример раздела;
- `cmd/<topic>/main.go` — запуск примера;
- `test/integration/lesson3_test.go` — проверка реального вывода программ.

`Example()` в каждом разделе использует большую часть функций. Поэтому после реализации задач нужно проверить не только unit-тесты, но и запуск команд.

## Порядок работы

1. Открой `docs/task.md` и выбери раздел.
2. Найди функции этого раздела в `internal/...`.
3. Реализуй `TODO` по порядку.
4. После каждой функции запускай тесты раздела.
5. Когда раздел готов, запусти соответствующий `cmd`.
6. В конце выполни полный `make ci`.

Не изменяй сигнатуры функций и не переноси учебную логику в `cmd/main.go`.

## Основные команды

```bash
make test-errors
make test-arrays
make test-structs
make test-newmake
make test-slices
make test-loops
make test-functions
make test-panics
make test-common
```

Запуск примеров:

```bash
make run-01_errors
make run-05_slices
make run-all
```

Финальная проверка:

```bash
make fmt
make test-unit
make test-integration
make ci
```
