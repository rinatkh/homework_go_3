package errorshomework

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestIsNil(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, true},
		{"simple", errors.New("x"), false},
		{"not found", ErrNotFound, false},
		{"wrapped not found", fmt.Errorf("repo: %w", ErrNotFound), false},
		{"field", &FieldError{Field: "name", Reason: "bad"}, false},
		{"wrapped field", fmt.Errorf("validate: %w", &FieldError{Field: "age", Reason: "bad"}), false},
		{"formatted", fmt.Errorf("failure %d", 1), false},
		{"empty text error", errors.New(""), false},
		{"double wrapped", fmt.Errorf("api: %w", fmt.Errorf("db: %w", ErrNotFound)), false},
		{"another", errors.New("another"), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNil(tc.err); got != tc.want {
				t.Fatalf("IsNil(%v)=%t, want %t", tc.err, got, tc.want)
			}
		})
	}
}

func TestRequireText(t *testing.T) {
	cases := []struct {
		name, field, value string
		wantErr            bool
	}{
		{"empty email", "email", "", true}, {"filled email", "email", "a@b", false},
		{"empty name", "name", "", true}, {"filled name", "name", "Maria", false},
		{"space is text", "comment", " ", false}, {"zero text", "code", "0", false},
		{"empty custom", "city", "", true}, {"unicode", "title", "Привет", false},
		{"emoji", "icon", "🙂", false}, {"empty field name", "", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := RequireText(tc.field, tc.value)
			if (err != nil) != tc.wantErr {
				t.Fatalf("RequireText(%q,%q) error=%v, wantErr=%t", tc.field, tc.value, err, tc.wantErr)
			}
			if tc.wantErr {
				var fieldErr *FieldError
				if !errors.As(err, &fieldErr) {
					t.Fatalf("expected *FieldError, got %T", err)
				}
				if fieldErr.Field != tc.field || fieldErr.Reason != "is required" {
					t.Fatalf("got %+v", fieldErr)
				}
			}
		})
	}
}

func TestValidatePositive(t *testing.T) {
	cases := []struct {
		name, field string
		value       int
		wantErr     bool
	}{
		{"one", "age", 1, false}, {"large", "amount", 1000, false}, {"zero", "age", 0, true},
		{"minus one", "count", -1, true}, {"negative", "price", -100, true}, {"two", "id", 2, false},
		{"max small", "score", 99, false}, {"empty field", "", 0, true}, {"min int sample", "offset", -9999, true},
		{"positive custom", "workers", 7, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePositive(tc.field, tc.value)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error=%v, wantErr=%t", err, tc.wantErr)
			}
			if tc.wantErr {
				var fieldErr *FieldError
				if !errors.As(err, &fieldErr) {
					t.Fatalf("expected FieldError")
				}
				if fieldErr.Field != tc.field || fieldErr.Reason != "must be positive" {
					t.Fatalf("got %+v", fieldErr)
				}
			}
		})
	}
}

func TestParseAge(t *testing.T) {
	cases := []struct {
		name, raw string
		want      int
		wantErr   bool
	}{
		{"positive", "26", 26, false}, {"zero", "0", 0, false}, {"negative", "-7", -7, false},
		{"plus", "+18", 18, false}, {"leading zero", "007", 7, false},
		{"letters", "abc", 0, true}, {"empty", "", 0, true}, {"decimal", "18.5", 0, true},
		{"spaces", " 20 ", 0, true}, {"mixed", "20x", 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseAge(tc.raw)
			if got != tc.want || (err != nil) != tc.wantErr {
				t.Fatalf("ParseAge(%q)=(%d,%v), want (%d, err=%t)", tc.raw, got, err, tc.want, tc.wantErr)
			}
			if tc.wantErr {
				if !strings.Contains(err.Error(), "parse age") {
					t.Fatalf("missing context: %v", err)
				}
				var numErr *strconv.NumError
				if !errors.As(err, &numErr) {
					t.Fatalf("source error is not preserved: %T", err)
				}
			}
		})
	}
}

func TestWrapOperation(t *testing.T) {
	baseA := errors.New("db down")
	baseB := errors.New("timeout")
	cases := []struct {
		name, operation string
		err             error
		wantNil         bool
		target          error
	}{
		{"nil save", "save", nil, true, nil}, {"nil empty operation", "", nil, true, nil},
		{"save", "save", baseA, false, baseA}, {"load", "load user", baseB, false, baseB},
		{"not found", "repository", ErrNotFound, false, ErrNotFound},
		{"wrapped", "service", fmt.Errorf("storage: %w", baseA), false, baseA},
		{"field", "validate", &FieldError{Field: "age", Reason: "bad"}, false, nil},
		{"empty op", "", baseB, false, baseB}, {"unicode", "загрузка", baseA, false, baseA},
		{"double", "handler", fmt.Errorf("usecase: %w", ErrNotFound), false, ErrNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := WrapOperation(tc.operation, tc.err)
			if tc.wantNil {
				if got != nil {
					t.Fatalf("got %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatal("got nil")
			}
			if tc.target != nil && !errors.Is(got, tc.target) {
				t.Fatalf("cause lost: %v", got)
			}
			if tc.operation != "" && !strings.Contains(got.Error(), tc.operation) {
				t.Fatalf("operation missing: %v", got)
			}
		})
	}
}

func TestNotFound(t *testing.T) {
	resources := []string{"user", "order", "product", "invoice", "profile", "address", "brand", "contract", "заказ", ""}
	for i, resource := range resources {
		t.Run(fmt.Sprintf("case_%02d", i+1), func(t *testing.T) {
			err := NotFound(resource)
			if err == nil {
				t.Fatal("got nil")
			}
			if !errors.Is(err, ErrNotFound) {
				t.Fatalf("ErrNotFound lost: %v", err)
			}
			if resource != "" && !strings.Contains(err.Error(), resource) {
				t.Fatalf("resource missing: %v", err)
			}
		})
	}
}

func TestIsNotFound(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false}, {"direct", ErrNotFound, true}, {"one wrap", fmt.Errorf("repo: %w", ErrNotFound), true},
		{"two wraps", fmt.Errorf("api: %w", fmt.Errorf("repo: %w", ErrNotFound)), true},
		{"text only", fmt.Errorf("repo: %v", ErrNotFound), false}, {"same text", errors.New("not found"), false},
		{"field", &FieldError{Field: "x", Reason: "not found"}, false}, {"other", errors.New("other"), false},
		{"resource helper", NotFound("user"), true}, {"wrapped helper", WrapOperation("load", NotFound("order")), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNotFound(tc.err); got != tc.want {
				t.Fatalf("got %t, want %t for %v", got, tc.want, tc.err)
			}
		})
	}
}

func TestSameCause(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	cases := []struct {
		name        string
		err, target error
		want        bool
	}{
		{"same direct", a, a, true}, {"different", a, b, false}, {"wrapped same", fmt.Errorf("x: %w", a), a, true},
		{"double wrapped", fmt.Errorf("y: %w", fmt.Errorf("x: %w", b)), b, true},
		{"text only", fmt.Errorf("x: %v", a), a, false}, {"nil nil", nil, nil, true},
		{"nil target", a, nil, false}, {"nil error", nil, a, false},
		{"not found", NotFound("user"), ErrNotFound, true}, {"same text different values", errors.New("a"), a, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SameCause(tc.err, tc.target); got != tc.want {
				t.Fatalf("got %t, want %t", got, tc.want)
			}
		})
	}
}

func TestFieldFrom(t *testing.T) {
	f1 := &FieldError{Field: "name", Reason: "blank"}
	f2 := &FieldError{Field: "age", Reason: "negative"}
	cases := []struct {
		name string
		err  error
		want *FieldError
		ok   bool
	}{
		{"direct name", f1, f1, true}, {"direct age", f2, f2, true}, {"wrapped", fmt.Errorf("validate: %w", f1), f1, true},
		{"double wrapped", fmt.Errorf("api: %w", fmt.Errorf("validate: %w", f2)), f2, true},
		{"plain", errors.New("plain"), nil, false}, {"nil", nil, nil, false},
		{"text field", errors.New("name: blank"), nil, false}, {"not found", ErrNotFound, nil, false},
		{"helper required", RequireText("email", ""), &FieldError{Field: "email", Reason: "is required"}, true},
		{"helper positive", ValidatePositive("count", 0), &FieldError{Field: "count", Reason: "must be positive"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := FieldFrom(tc.err)
			if ok != tc.ok {
				t.Fatalf("ok=%t, want %t", ok, tc.ok)
			}
			if !tc.ok {
				if got != nil {
					t.Fatalf("got %+v, want nil", got)
				}
				return
			}
			if got == nil || got.Field != tc.want.Field || got.Reason != tc.want.Reason {
				t.Fatalf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestFieldName(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"direct", &FieldError{Field: "name", Reason: "bad"}, "name"},
		{"wrapped", fmt.Errorf("validate: %w", &FieldError{Field: "email", Reason: "bad"}), "email"},
		{"double", fmt.Errorf("api: %w", fmt.Errorf("service: %w", &FieldError{Field: "age", Reason: "bad"})), "age"},
		{"empty field", &FieldError{Field: "", Reason: "bad"}, ""}, {"plain", errors.New("x"), ""},
		{"nil", nil, ""}, {"not found", ErrNotFound, ""}, {"required", RequireText("city", ""), "city"},
		{"positive", ValidatePositive("price", -1), "price"}, {"text only", errors.New("field: name"), ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FieldName(tc.err); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestFirstError(t *testing.T) {
	a := errors.New("a")
	b := errors.New("b")
	cases := []struct {
		name string
		errs []error
		want error
	}{
		{"nil slice", nil, nil}, {"empty", []error{}, nil}, {"only nil", []error{nil}, nil},
		{"first", []error{a, b}, a}, {"after nil", []error{nil, a, b}, a}, {"second after nils", []error{nil, nil, b}, b},
		{"field", []error{nil, &FieldError{Field: "x", Reason: "bad"}}, nil},
		{"not found", []error{nil, ErrNotFound}, ErrNotFound}, {"last", []error{nil, nil, nil, a}, a},
		{"many", []error{nil, b, a, ErrNotFound}, b},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FirstError(tc.errs)
			if tc.name == "field" {
				var field *FieldError
				if !errors.As(got, &field) {
					t.Fatalf("got %v", got)
				}
				return
			}
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCountErrors(t *testing.T) {
	e := errors.New("x")
	cases := []struct {
		name string
		errs []error
		want int
	}{
		{"nil", nil, 0}, {"empty", []error{}, 0}, {"one nil", []error{nil}, 0}, {"one error", []error{e}, 1},
		{"mixed", []error{nil, e, nil}, 1}, {"two", []error{e, ErrNotFound}, 2},
		{"three with nil", []error{e, nil, ErrNotFound, &FieldError{}}, 3}, {"all nil", []error{nil, nil, nil}, 0},
		{"five", []error{e, e, e, e, e}, 5}, {"wrapped", []error{fmt.Errorf("x: %w", e), nil}, 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountErrors(tc.errs); got != tc.want {
				t.Fatalf("got %d, want %d", got, tc.want)
			}
		})
	}
}

func TestErrorText(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, ""}, {"simple", errors.New("boom"), "boom"}, {"empty", errors.New(""), ""},
		{"not found", ErrNotFound, "not found"}, {"field", &FieldError{Field: "age", Reason: "bad"}, "age: bad"},
		{"wrapped", fmt.Errorf("load: %w", ErrNotFound), "load: not found"}, {"number", fmt.Errorf("code %d", 42), "code 42"},
		{"unicode", errors.New("ошибка"), "ошибка"}, {"required", RequireText("email", ""), "email: is required"},
		{"positive", ValidatePositive("count", 0), "count: must be positive"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ErrorText(tc.err); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestValidateUser(t *testing.T) {
	cases := []struct {
		name, userName, email string
		age                   int
		wantField             string
		wantNil               bool
	}{
		{"valid", "Maria", "m@example.com", 20, "", true}, {"valid one", "A", "a", 1, "", true},
		{"missing name", "", "a", 20, "name", false}, {"missing email", "Maria", "", 20, "email", false},
		{"zero age", "Maria", "a", 0, "age", false}, {"negative age", "Maria", "a", -1, "age", false},
		{"all invalid first name", "", "", 0, "name", false}, {"email and age first email", "Maria", "", 0, "email", false},
		{"space values", " ", " ", 2, "", true}, {"unicode", "Мария", "почта", 30, "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateUser(tc.userName, tc.email, tc.age)
			if tc.wantNil {
				if err != nil {
					t.Fatalf("got %v", err)
				}
				return
			}
			var fieldErr *FieldError
			if !errors.As(err, &fieldErr) {
				t.Fatalf("expected FieldError, got %T", err)
			}
			if fieldErr.Field != tc.wantField {
				t.Fatalf("field=%q, want %q", fieldErr.Field, tc.wantField)
			}
		})
	}
}

func TestClassify(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{"nil", nil, "none"}, {"not found direct", ErrNotFound, "not_found"},
		{"not found wrapped", fmt.Errorf("repo: %w", ErrNotFound), "not_found"},
		{"not found double", WrapOperation("load", NotFound("user")), "not_found"},
		{"field name", &FieldError{Field: "name", Reason: "bad"}, "field:name"},
		{"field age wrapped", fmt.Errorf("validate: %w", &FieldError{Field: "age", Reason: "bad"}), "field:age"},
		{"field empty", &FieldError{Field: "", Reason: "bad"}, "field:"},
		{"plain", errors.New("x"), "other"}, {"same text", errors.New("not found"), "other"},
		{"text field", errors.New("name: bad"), "other"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Classify(tc.err); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestExampleOutput(t *testing.T) {
	want := "nil=true required=email: is required age=26 parseErr=false notFound=true field=email count=2 first=email: is required class=not_found"
	if got := Example(); got != want {
		t.Fatalf("Example()=%q, want %q", got, want)
	}
}
