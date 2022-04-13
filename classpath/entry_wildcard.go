package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// wildcardEntry 实际上也是 compositeEntry
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	var compositeEntry []Entry
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		return nil
	}
	return compositeEntry
}
