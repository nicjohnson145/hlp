package hlp

import (
	"fmt"
)

func ExampleInvert() {
	fmt.Println(Invert(map[string]int{"a": 1}))
	// Output: map[1:a]
}
