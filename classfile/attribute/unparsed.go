package attribute

import (
	"pineapple/classfile"
)

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

// 没有在 Java 虚拟机规范中定义的属性，通过此方法读取属性
func (u *UnparsedAttribute) readInfo(reader *classfile.ClassReader) {
	u.info = reader.ReadByte(u.length)
}
