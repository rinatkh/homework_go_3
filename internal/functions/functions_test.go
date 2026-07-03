package functions

import "testing"

func TestFullName(t *testing.T) {
	if got := FullName("Мария", "Иванова"); got != "Мария Иванова" {
		t.Fatalf("FullName=%q", got)
	}
}

func TestPriceWithDiscount(t *testing.T) {
	if got := PriceWithDiscount(1000, 15); got != 850 {
		t.Fatalf("PriceWithDiscount=%v", got)
	}
	if got := PriceWithDiscount(1000, 150); got != 0 {
		t.Fatalf("PriceWithDiscount over 100=%v", got)
	}
}

func TestSafeDivide(t *testing.T) {
	got, ok := SafeDivide(17, 5)
	if got != 3 || !ok {
		t.Fatalf("SafeDivide=%d,%t", got, ok)
	}
	_, ok = SafeDivide(1, 0)
	if ok {
		t.Fatalf("SafeDivide by zero ok=true")
	}
}

func TestMinMax(t *testing.T) {
	minValue, maxValue := MinMax(10, -4)
	if minValue != -4 || maxValue != 10 {
		t.Fatalf("MinMax=%d,%d", minValue, maxValue)
	}
}

func TestNormalizeEmail(t *testing.T) {
	if got := NormalizeEmail("  USER@EXAMPLE.COM "); got != "user@example.com" {
		t.Fatalf("NormalizeEmail=%q", got)
	}
}

func TestFormatUser(t *testing.T) {
	if got := FormatUser(7, "Maria", true); got != "#7 Maria (active)" {
		t.Fatalf("FormatUser=%q", got)
	}
}

func TestApplyOperation(t *testing.T) {
	got := ApplyOperation(3, 4, func(a, b int) int { return a * b })
	if got != 12 {
		t.Fatalf("ApplyOperation=%d", got)
	}
	if got := ApplyOperation(3, 4, nil); got != 0 {
		t.Fatalf("ApplyOperation nil=%d", got)
	}
}

func TestSumVariadic(t *testing.T) {
	if got := SumVariadic(1, 2, 3, 4); got != 10 {
		t.Fatalf("SumVariadic=%d", got)
	}
}

func TestBuildGreeting(t *testing.T) {
	if got := BuildGreeting("ru", "Мария"); got != "Привет, Мария" {
		t.Fatalf("BuildGreeting=%q", got)
	}
}

func TestValidatePassword(t *testing.T) {
	if ok, _ := ValidatePassword("abc"); ok {
		t.Fatalf("short password ok=true")
	}
	if ok, _ := ValidatePassword("abcdefgh"); ok {
		t.Fatalf("password without digit ok=true")
	}
	if ok, msg := ValidatePassword("abc12345"); !ok || msg != "ok" {
		t.Fatalf("ValidatePassword=%t,%q", ok, msg)
	}
}

func TestSplitFullName(t *testing.T) {
	first, last := SplitFullName("Мария Иванова")
	if first != "Мария" || last != "Иванова" {
		t.Fatalf("SplitFullName=%q,%q", first, last)
	}
}

func TestCalcOrderTotal(t *testing.T) {
	if got := CalcOrderTotal(1000, 150, 10); got != 1050 {
		t.Fatalf("CalcOrderTotal=%v", got)
	}
}

func TestMakeCounter(t *testing.T) {
	counter := MakeCounter(10)
	if counter() != 11 || counter() != 12 {
		t.Fatalf("counter does not keep state")
	}
}

func TestSwap(t *testing.T) {
	left, right := Swap("a", "b")
	if left != "b" || right != "a" {
		t.Fatalf("Swap=%q,%q", left, right)
	}
}

func TestAverage(t *testing.T) {
	avg, ok := Average(2, 4, 6)
	if avg != 4 || !ok {
		t.Fatalf("Average=%v,%t", avg, ok)
	}
	_, ok = Average()
	if ok {
		t.Fatalf("Average empty ok=true")
	}
}
