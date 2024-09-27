package hlp

import (
	"fmt"
	"regexp"
)

func ExampleExtractNamedMatches() {
	exp := regexp.MustCompile(`(?P<first>\d+)\.(?P<second>\d+)\.(\d+)`)
	got := ExtractNamedMatches(exp, exp.FindStringSubmatch("192.168.1"))
	fmt.Println(got)
	// Output: map[first:192 second:168]
}
