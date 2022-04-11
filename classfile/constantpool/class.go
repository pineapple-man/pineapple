package constantpool

import (
	"pineapple/classfile"
)

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/

type ConstantClassInfo struct {
	constantPool ConstantPool
	// 全限定名常量的索引
	nameIndex uint16
}

func (c *ConstantClassInfo) readInfo(reader *classfile.ClassReader) {
	c.nameIndex = reader.ReadUint16()
}
func (c *ConstantClassInfo) Name() string {
	return c.constantPool.GetUtf8(c.nameIndex)
}
