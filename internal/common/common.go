package common

import (
	"errors"
	"fmt"
)

var ErrNoActiveUsers = errors.New("no active users")

type User struct {
	ID     int
	Name   string
	Active bool
}

type Order struct {
	ID     int
	Amount int
	Paid   bool
}

type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

// 01. ActiveUserReport формирует отчёт по активным пользователям.
//
// TODO: сохраните порядок пользователей. Формат результата:
// "active=<количество> names=<имя1>,<имя2>".
// Если активных пользователей нет, верните "" и ErrNoActiveUsers.
func ActiveUserReport(users []User) (string, error) { return "", nil }

// 02. DiscountOrder возвращает копию заказа со скидкой.
//
// TODO: Amount не может быть отрицательным, percent должен быть от 0 до 100.
// Для некорректного поля верните *FieldError с Field "amount" или "percent".
// Исходный order не изменяйте. Вычисления выполняются целыми числами.
func DiscountOrder(order Order, percent int) (Order, error) { return Order{}, nil }

// 03. AppendToIsolatedWindow создаёт окно items[low:high] с ограниченной
// capacity и добавляет values.
//
// TODO: допустимы границы 0 <= low <= high <= len(items).
// При неверных границах верните ошибку и nil-слайсы.
// При корректных границах исходный items не должен измениться из-за append.
func AppendToIsolatedWindow(items []int, low, high int, values ...int) (source, window []int, err error) {
	return nil, nil, nil
}
