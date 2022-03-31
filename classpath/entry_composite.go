package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (c CompositeEntry) ReadClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.ReadClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	var compositeEntryStrings = make([]string, len(c))
	for i, entry := range c {
		compositeEntryStrings[i] = entry.String()
	}
	return strings.Join(compositeEntryStrings, pathListSeparator)
}
