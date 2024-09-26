package hlp

import (
	"fmt"
)

func ExampleExtractRange() {
	got, _ := ExtractRange([]string{"apple", "banana", "cherry", "date", "egg", "fries", "grapes"}, "0,2,5-")
	fmt.Println(got)
	// Output: [apple cherry fries grapes]
}
