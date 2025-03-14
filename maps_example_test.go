package hlp

import (
	"fmt"
)

func ExampleInvert() {
	fmt.Println(Invert(map[string]int{"a": 1}))
	// Output: map[1:a]
}

func ExampleFilteredSliceFromMap() {
	input := map[string]string{
		"a": "one",
		"b": "two",
		"c": "three",
	}
	got := FilteredSliceFromMap(input, func(key string, value string) (string, bool) {
		return key + "|" + value, key != "c"
	})
	fmt.Println(got)
	// Output: [a|one b|two]
}

func ExampleSliceFromMap() {
	input := map[string]string{
		"a": "one",
		"b": "two",
		"c": "three",
	}
	got := SliceFromMap(input, func(key string, value string) string {
		return key + "|" + value
	})
	fmt.Println(got)
	// Output: [a|one b|two c|three]
}
