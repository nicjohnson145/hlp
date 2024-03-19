package testhlp

import (
	"github.com/stretchr/testify/require"
)

type TestingT interface {
	require.TestingT
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Fail()
}
