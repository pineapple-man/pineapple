package classpath

import (
	"os"
	"strings"
)

const PATH_LIST_SEPARATOR = string(os.PathListSeparator)

// Entry 类路径项
type Entry interface {
	// ReadClass 方法的参数是 class 文件的相对路径，路径之间用斜线 `/` 分隔，文件名有 `.class` 后缀
	ReadClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, PATH_LIST_SEPARATOR) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return NewDirEntry(path)
	}
	return NewDirEntry(path)
}
