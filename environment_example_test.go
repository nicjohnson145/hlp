package hlp

import (
	"fmt"
)

func ExampleDefaultEnv() {
	fmt.Println(DefaultEnv("FOO", "BAR"))
	// Output: BAR
}
