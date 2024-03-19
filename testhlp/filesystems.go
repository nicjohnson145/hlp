package testhlp

import (
	"io/fs"
	"sort"

	"github.com/google/go-cmp/cmp"
	"github.com/nicjohnson145/hlp"
	"github.com/nicjohnson145/hlp/set"
	"github.com/stretchr/testify/require"
)

type FSItem struct {
	IsDir bool
	Content string
}

func (f *FSItem) TypeString() string {
	if f.IsDir {
		return "directory"
	}
	return "file"
}

func CompareFS(t TestingT, expected fs.FS, actual fs.FS) {
	t.Helper()

	expectedContents := FSContents(t, expected)
	actualContents := FSContents(t, actual)

	expectedPaths := set.New(hlp.Keys(expectedContents)...)
	actualPaths := set.New(hlp.Keys(actualContents)...)

	missingPaths := expectedPaths.Difference(actualPaths).AsSlice()
	if len(missingPaths) > 0 {
		t.Log("the following paths are expected but not present")
		for _, path := range missingPaths {
			t.Logf("* %v", path)
		}
		t.Fail()
	}

	extraPaths := actualPaths.Difference(expectedPaths).AsSlice()
	if len(extraPaths) > 0 {
		t.Log("the following paths are in the output, but not expected to be")
		for _, path := range extraPaths {
			t.Logf("* %v", path)
		}
		t.Fail()
	}

	sharedPaths := expectedPaths.Intersection(actualPaths).AsSlice()
	// Make unit testings stable
	sort.Strings(sharedPaths)
	for _, path := range sharedPaths {
		expectedContent := expectedContents[path]
		actualContent := actualContents[path]

		if expectedContent.IsDir != actualContent.IsDir {
			t.Logf("path %v expected to be a %v but got %v", path, expectedContent.TypeString(), actualContent.TypeString())
			t.Fail()
			continue
		}

		if !expectedContent.IsDir {
			if diff := cmp.Diff(expectedContent.Content, actualContent.Content); diff != "" {
				t.Logf("Content mismatch at %v (-want, +got): \n%v", path, diff)
				t.Fail()
			}
		}
	}
}

func FSContents(t TestingT, fsys fs.FS) map[string]FSItem {
	t.Helper()

	contents := map[string]FSItem{}

	require.NoError(t, fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var content string
		if !d.IsDir() {
			contentBytes, err := fs.ReadFile(fsys, path)
			require.NoError(t, err, "error reading ", path)
			content = string(contentBytes)
		}

		contents[path] = FSItem{
			IsDir: d.IsDir(),
			Content: content,
		}

		return nil
	}))

	return contents
}
