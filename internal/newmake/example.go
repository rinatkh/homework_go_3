package newmake

import "fmt"

func intValue(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}

func stringValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func Example() string {
	number := NewInt(26)
	text := NewString("go")
	ints := MakeIntsWithCapacity(2, 5)
	appended := MakeAndAppend(4, 10, 20, 30)
	length, capacity, isNil := SliceState(ints)
	_, _, emptyNil := EmptySliceState()
	user := MakeUserScores("Maria", 3)
	slicePointer := NewSlicePointer()
	channel := MakeBoolChannel(2)
	defer close(channel)

	userName := ""
	scoreCount := 0
	if user != nil {
		userName = user.Name
		scoreCount = len(user.Scores)
	}

	pointedSliceNil := false
	if slicePointer != nil {
		pointedSliceNil = *slicePointer == nil
	}

	return fmt.Sprintf(
		"int=%d string=%s slice=%v len=%d cap=%d nil=%t appended=%v emptyNil=%t user=%s/%d newSliceNil=%t mapNil=%t chanCap=%d",
		intValue(number),
		stringValue(text),
		ints,
		length,
		capacity,
		isNil,
		appended,
		emptyNil,
		userName,
		scoreCount,
		pointedSliceNil,
		MakeStringIntMap() == nil,
		cap(channel),
	)
}
