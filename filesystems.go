package hlp

import (
	"io"
	"io/fs"

	"github.com/jarxorg/wfs"
)

func CopyFS(dest wfs.WriteFileFS, src fs.FS, root string) error {
	return fs.WalkDir(src, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d == nil {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return err
		}

		if d.IsDir() {
			return dest.MkdirAll(path, info.Mode())
		}
		srcFile, err := src.Open(path)
		if err != nil {
			return err
		}
		destFile, err := dest.CreateFile(path, info.Mode())
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		return err
	})
}
