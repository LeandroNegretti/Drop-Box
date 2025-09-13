package utils

import (
	"io"
	"os"
	"path/filepath"
)

func SaveFile(destDir, filename string, src io.Reader) error {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	dstPath := filepath.Join(destDir, filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
