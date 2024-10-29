package hlp

import (
	"regexp"
)

// ExtractNamedMatches converts the array returned by `regexp.FindStringSubmatch` into a map, whose keys are the capture
// group names, and whose values are the results
func ExtractNamedMatches(exp *regexp.Regexp, submatches []string) map[string]string {
	result := map[string]string{}
	for i, name := range exp.SubexpNames() {
		if i != 0 && name != "" && i < len(submatches){
			result[name] = submatches[i]
		}
	}

	return result
}
