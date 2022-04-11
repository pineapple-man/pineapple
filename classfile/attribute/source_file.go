package attribute

import (
	"pineapple/classfile"
	"pineapple/classfile/constantpool"
)

/*
SourceFile 是可选定长属性，只会出现在 ClassFile 结构中，用于指出源文件名，
attribute_length 必须是 2，sourcefile_index是常量池索引
SourceFile_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 sourcefile_index;
}
*/

type SourceFileAttribute struct {
	constantPool    constantpool.ConstantPool
	sourceFileIndex uint16
}

func (s *SourceFileAttribute) readInfo(reader *classfile.ClassReader) {
	s.sourceFileIndex = reader.ReadUint16()
}

// FileName 读取文件名称
func (s *SourceFileAttribute) FileName() string {
	return s.constantPool.GetUtf8(s.sourceFileIndex)
}
