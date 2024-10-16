package hlp

import (
	"os"
)

// DefaultEnv tries to look up an environment variable, returning the value if set, or defaultVal if not set
func DefaultEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}
