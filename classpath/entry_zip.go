package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry zip 或 jar 文件形式的类路径
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}
func (z *ZipEntry) ReadClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer func(reader *zip.ReadCloser) {
		_ = reader.Close()
	}(reader)

	for _, f := range reader.File {
		if f.Name == className {
			open, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer func(open io.ReadCloser) {
				err := open.Close()
				if err != nil {
					panic(err)
				}
			}(open)
			data, err := ioutil.ReadAll(open)
			if err != nil {
				return nil, nil, err
			}
			return data, z, err
		}
	}
	return nil, nil, errors.New("class not fount: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
