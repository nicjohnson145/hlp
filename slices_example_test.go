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

func ExampleMax() {
	fmt.Println(Max([]int{1, 3, 2, 0}))
	// Output: 3
}

func ExampleMaxBy() {
	fmt.Println(MaxBy([]int{1, 3, 2, 0}, func(item int, high int) bool { return item > high }))
	// Output: 3
}

func ExampleMin() {
	fmt.Println(Min([]int{1, 3, 2, 0}))
	// Output: 0
}

func ExampleMinBy() {
	fmt.Println(MinBy([]int{1, 3, 2, 0}, func(item int, low int) bool { return item < low }))
	// Output: 0
}
