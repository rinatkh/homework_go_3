package panics

import "fmt"

func Example() string {
	normalPanicked := SafeRun(func() {})
	boomPanicked := SafeRun(func() { panic("boom") })
	message := RecoverMessage(func() { panic("broken") })
	parsed, parsedOK := TryParse("42")
	invalid, invalidOK := TryParse("x")
	indexed, indexOK := SafeIndex([]int{10, 20, 30}, 1)
	missing, missingOK := SafeIndex([]int{10}, 4)
	panicErr := PanicToError(func() { panic("bad state") })
	completed, panicked := RunSequenceSafe([]func(){func() {}, func() { panic("x") }, func() {}})
	number := 7

	return fmt.Sprintf(
		"normal=%t boom=%t message=%s parse=%d/%t invalid=%d/%t index=%d/%t missing=%d/%t error=%s events=%v outside=%t sequence=%d/%d pointer=%d",
		normalPanicked,
		boomPanicked,
		message,
		parsed,
		parsedOK,
		invalid,
		invalidOK,
		indexed,
		indexOK,
		missing,
		missingOK,
		panicErr,
		DeferBeforePanic(),
		RecoverOutsideDefer(),
		completed,
		panicked,
		MustNonNil(&number),
	)
}
