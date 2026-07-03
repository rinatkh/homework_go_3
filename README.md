# homework_go_3: циклы, функции, массивы, слайсы

Домашняя работа к третьему занятию курса Go.

## Что нужно сделать

Реализовать все функции с `TODO` внутри `internal/*`.

Структура специально сделана как в production-проекте:

- `cmd/*/main.go` только вызывает `Example()`;
- вся логика находится в `internal/*`;
- unit-тесты лежат рядом с кодом в `internal/*`;
- integration-тесты проверяют запуск `cmd/*` через `go run`;
- `Makefile` содержит команды для локальной проверки;
- `.github/workflows/ci.yml` запускает проверки в GitHub Actions.

## Команды

```bash
make test-unit
make test-integration
make run-all
make ci
```

Если тесты падают — это нормально в начале выполнения. Нужно реализовать TODO до зелёного CI.
