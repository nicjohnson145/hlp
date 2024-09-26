package hlp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidRangeExpressionError = errors.New("invalid range expression")
	ErrOpenRangeNotAtEndError      = errors.New("open range can only be final range specifed")
	ErrRangeExceedLengthError      = errors.New("range exceeds slice length")
	ErrUnableToParseError          = errors.New("unable to parse range expression")
)

// ParseRange parses the given range expression within the context of the supplied total length. It returns a list of
// indexes that correspond to the range expression or error
func ParseRange(totalLength int, expr string) ([]int, error) {
	indexes := []int{}

	sections := strings.Split(expr, ",")

	for i, section := range sections {
		// Is it a range expression
		if strings.Contains(section, "-") {
			parts := strings.Split(section, "-")
			if len(parts) > 2 {
				return nil, fmt.Errorf("%w: %w: subranges can contain at most 2 elements",ErrInvalidRangeExpressionError, ErrUnableToParseError)
			}
			parts = Filter(parts, func(s string, _ int) bool { return s != "" })

			var start int
			var end int
			// Is it an open ended range? If so, it better be the last one
			if len(parts) == 1 {
				if i != len(sections) - 1{
					return nil, fmt.Errorf("%w: %w", ErrInvalidRangeExpressionError, ErrOpenRangeNotAtEndError)
				}
				s, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, fmt.Errorf("%w: %w: %w", ErrInvalidRangeExpressionError, ErrUnableToParseError, err)
				}
				start = s
				end = totalLength - 1
			} else { // otherwise, parse both of them
				s, err := strconv.Atoi(parts[0])
				if err != nil {
					return nil, fmt.Errorf("%w: %w: %w", ErrInvalidRangeExpressionError, ErrUnableToParseError, err)
				}
				start = s
				e, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("%w: %w: %w", ErrInvalidRangeExpressionError, ErrUnableToParseError, err)
				}
				end = e
			}
			for i := start; i <= end; i++ {
				indexes = append(indexes, i)
			}
		} else { // otherwise treat its a single index
			idx, err := strconv.Atoi(section)
			if err != nil {
				return nil, fmt.Errorf("%w: %w: %w", ErrInvalidRangeExpressionError, ErrUnableToParseError, err)
			}
			indexes = append(indexes, idx)
		}
	}

	return indexes, nil
}
