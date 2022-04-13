package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 目录形式的类路径
type DirEntry struct {
	AbsDir string
}

func NewDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 如果转换过程出现错误，终止程序执行
		panic(err)
	}
	return &DirEntry{AbsDir: absDir}
}

func (e *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.AbsDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, e, err
}

func (e *DirEntry) String() string {
	return e.AbsDir
}
