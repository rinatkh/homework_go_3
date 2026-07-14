package errorshomework

import (
	"errors"
	"fmt"
	"strings"
)

func Example() string {
	parsed, parseErr := ParseAge("26")
	requiredErr := RequireText("email", "")
	wrapped := WrapOperation("load user", NotFound("user"))
	field, fieldOK := FieldFrom(requiredErr)
	first := FirstError([]error{nil, requiredErr, errors.New("later")})

	fieldName := ""
	if fieldOK {
		fieldName = field.Field
	}

	return strings.Join([]string{
		fmt.Sprintf("nil=%t", IsNil(nil)),
		fmt.Sprintf("required=%s", ErrorText(requiredErr)),
		fmt.Sprintf("age=%d parseErr=%t", parsed, parseErr != nil),
		fmt.Sprintf("notFound=%t", IsNotFound(wrapped)),
		fmt.Sprintf("field=%s", fieldName),
		fmt.Sprintf("count=%d", CountErrors([]error{nil, requiredErr, wrapped})),
		fmt.Sprintf("first=%s", ErrorText(first)),
		fmt.Sprintf("class=%s", Classify(wrapped)),
	}, " ")
}
