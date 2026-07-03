GO ?= go
COVERAGE_FILE ?= coverage.out
COVERAGE_MIN ?= 80
PACKAGE_NAME ?= homework-go-3-linux-amd64

.PHONY: help deps-check mod-check fmt fmt-check vet compile test test-unit test-integration test-race coverage coverage-check build package clean run-all ci \
run-loops run-functions run-arrays run-slices \
test-loops test-functions test-arrays test-slices

help: ## Показать список команд
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "%-24s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps-check: ## Скачать и проверить зависимости Go-модулей
	$(GO) mod download
	$(GO) mod verify

mod-check: ## Проверить, что go.mod/go.sum не меняются после go mod tidy
	$(GO) mod tidy
	@if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then \
		if [ -n "$$($(GO) env GOMOD >/dev/null; git status --short -- go.mod go.sum)" ]; then \
			git status --short -- go.mod go.sum; \
			echo "go.mod/go.sum changed after go mod tidy"; \
			exit 1; \
		fi; \
	else \
		echo "not a git repository, skip go.mod/go.sum diff check"; \
	fi

fmt: ## Отформатировать Go-файлы
	$(GO) fmt ./...

fmt-check: ## Проверить gofmt без изменения файлов
	@test -z "$$(gofmt -l . | grep -v '^bin/')" || (echo "Run gofmt for files above" && gofmt -l . && exit 1)

vet: ## Запустить go vet
	$(GO) vet ./...

compile: ## Проверить, что все пакеты компилируются без запуска тестов
	$(GO) test -run '^$$' ./...

test: test-unit test-integration ## Запустить все тесты

test-unit: ## Запустить unit-тесты по internal-пакетам
	$(GO) test ./internal/...

test-integration: ## Запустить integration-тесты по cmd/main
	$(GO) test ./test/integration

test-race: ## Запустить тесты с race detector
	$(GO) test -race ./internal/...

coverage: ## Посчитать coverage по internal-пакетам
	$(GO) test ./internal/... -coverprofile=$(COVERAGE_FILE)
	$(GO) tool cover -func=$(COVERAGE_FILE)

coverage-check: coverage ## Проверить минимальный coverage
	@total=$$($(GO) tool cover -func=$(COVERAGE_FILE) | awk '/total:/ {gsub("%", "", $$3); print $$3}'); \
	awk -v total="$$total" -v min="$(COVERAGE_MIN)" 'BEGIN { if (total + 0 < min + 0) { printf("coverage %.1f%% is below %.1f%%\n", total, min); exit 1 } else { printf("coverage %.1f%% is OK\n", total) } }'

build: ## Собрать все main в bin/
	mkdir -p bin
	$(GO) build -o bin/01_loops ./cmd/01_loops
	$(GO) build -o bin/02_functions ./cmd/02_functions
	$(GO) build -o bin/03_arrays ./cmd/03_arrays
	$(GO) build -o bin/04_slices ./cmd/04_slices

package: build ## Упаковать собранные бинарники в tar.gz
	tar -czf bin/$(PACKAGE_NAME).tar.gz \
		-C bin \
		01_loops \
		02_functions \
		03_arrays \
		04_slices

clean: ## Удалить временные файлы
	rm -rf bin $(COVERAGE_FILE)

run-all: ## Запустить все main по очереди
	$(GO) run ./cmd/01_loops
	$(GO) run ./cmd/02_functions
	$(GO) run ./cmd/03_arrays
	$(GO) run ./cmd/04_slices

run-loops: ## Запустить 01_loops
	$(GO) run ./cmd/01_loops

run-functions: ## Запустить 02_functions
	$(GO) run ./cmd/02_functions

run-arrays: ## Запустить 03_arrays
	$(GO) run ./cmd/03_arrays

run-slices: ## Запустить 04_slices
	$(GO) run ./cmd/04_slices

test-loops: ## Unit-тесты пакета internal/loops
	$(GO) test ./internal/loops

test-functions: ## Unit-тесты пакета internal/functions
	$(GO) test ./internal/functions

test-arrays: ## Unit-тесты пакета internal/arrays
	$(GO) test ./internal/arrays

test-slices: ## Unit-тесты пакета internal/slices
	$(GO) test ./internal/slices

ci: deps-check mod-check fmt-check vet test-unit test-integration test-race coverage-check build package ## Полная локальная проверка как в CI
