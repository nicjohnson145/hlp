package hlp

import (
	"fmt"
)

func ExampleExtractRange() {
	got, _ := ExtractRange([]string{"apple", "banana", "cherry", "date", "egg", "fries", "grapes"}, "0,2,5-")
	fmt.Println(got)
	// Output: [apple cherry fries grapes]
}

func ExampleAny() {
	fmt.Println(Any([]string{"ab", "cde", "fg"}, func(x string) bool {
		return len(x) > 2
	}))
	// Output: true
}

func ExampleAll() {
	fmt.Println(All([]string{"ab", "cde", "fg"}, func(x string) bool {
		return len(x) > 2
	}))
	// Output: false
}

func ExampleFirst() {
	fmt.Println(First([]int{1, 3, 4, 6}, func(x int) bool {
		return x%2 == 0
	}))
	// Output: 2
}
