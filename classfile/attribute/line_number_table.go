package attribute

import (
	"pineapple/classfile"
)

/*

LineNumberTable_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 line_number_table_length;
	{
		u2 start_pc;
		u2 line_number;
	} line_number_table[line_number_table_length];
}
*/

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (l *LineNumberTableAttribute) readInfo(reader *classfile.ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	lineNumberTableEntries := make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lineNumberTableEntries {
		lineNumberTableEntries[i] = &LineNumberTableEntry{
			startPc:    reader.ReadUint16(),
			lineNumber: reader.ReadUint16(),
		}
	}
}
