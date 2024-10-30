package testhlp

import (
	"strings"
	"testing"

	"github.com/psanford/memfs"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCompareFS(t *testing.T) {
	aFS := memfs.New()
	require.NoError(t, aFS.MkdirAll("different", 0775))
	require.NoError(t, aFS.WriteFile("different/differentfile.txt", []byte("dir_a\n"), 0774))
	require.NoError(t, aFS.MkdirAll("directory_in_a", 0775))
	require.NoError(t, aFS.MkdirAll("same", 0775))
	require.NoError(t, aFS.WriteFile("same/samefile.txt", []byte(`same-content`), 0774))
	require.NoError(t, aFS.WriteFile("file_in_a", []byte(`content`), 0774))
	require.NoError(t, aFS.WriteFile("only_in_a", []byte(`content`), 0774))

	bFS := memfs.New()
	require.NoError(t, bFS.MkdirAll("different", 0775))
	require.NoError(t, bFS.WriteFile("different/differentfile.txt", []byte("dir_b\n"), 0774))
	require.NoError(t, bFS.WriteFile("directory_in_a", []byte(`content`), 0774))
	require.NoError(t, bFS.MkdirAll("same", 0775))
	require.NoError(t, bFS.WriteFile("same/samefile.txt", []byte(`same-content`), 0774))
	require.NoError(t, bFS.MkdirAll("file_in_a", 0775))
	require.NoError(t, bFS.WriteFile("only_in_b", []byte(`content`), 0774))

	testT := &MockTestingT{}
	testT.EXPECT().Helper().Return()
	testT.EXPECT().Log(mock.Anything).Return()
	testT.EXPECT().Logf(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return()
	testT.EXPECT().Logf(mock.Anything, mock.Anything, mock.Anything).Return()
	testT.EXPECT().Logf(mock.Anything, mock.Anything).Return()
	testT.EXPECT().Fail().Return()

	CompareFS(testT, aFS, bFS)

	type call struct {
		Method    string
		Arguments []any
	}
	// for each call, just extract the method & the arguments
	calls := []call{}
	for _, c := range testT.Mock.Calls {
		if c.Method == "Helper" {
			continue
		}
		var newArgs []any
		for _, a := range c.Arguments {
			strVal, ok := a.(string)
			if !ok {
				newArgs = append(newArgs, a)
				continue
			}

			newArgs = append(newArgs, strings.ReplaceAll(strVal, "\u00a0", " "))
		}
		calls = append(calls, call{
			Method:    c.Method,
			Arguments: newArgs,
		})
	}

	expectedCalls := []call{
		{
			Method:    "Log",
			Arguments: []any{"the following paths are expected but not present"},
		},
		{
			Method:    "Logf",
			Arguments: []any{"* %v", "only_in_a"},
		},
		{
			Method: "Fail",
		},
		{
			Method:    "Log",
			Arguments: []any{"the following paths are in the output, but not expected to be"},
		},
		{
			Method:    "Logf",
			Arguments: []any{"* %v", "only_in_b"},
		},
		{
			Method: "Fail",
		},
		{
			Method:    "Logf",
			Arguments: []any{"Content mismatch at %v (-want, +got): \n%v", "different/differentfile.txt", "  string(\n- \t\"dir_a\\n\",\n+ \t\"dir_b\\n\",\n  )\n"},
		},
		{
			Method: "Fail",
		},
		{
			Method:    "Logf",
			Arguments: []any{"path %v expected to be a %v but got %v", "directory_in_a", "directory", "file"},
		},
		{
			Method: "Fail",
		},
		{
			Method:    "Logf",
			Arguments: []any{"path %v expected to be a %v but got %v", "file_in_a", "file", "directory"},
		},
		{
			Method: "Fail",
		},
	}

	require.Equal(t, expectedCalls, calls)

}
