package classfile

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/

type ConstantStringInfo struct {
	constantPool ConstantPool
	stringIndex  uint16
}

// 按照索引从常量池中查找字符串
func (c *ConstantStringInfo) String() string {
	return c.constantPool.GetUtf8(c.stringIndex)
}

// 读取常量池索引
func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.stringIndex = reader.readUint16()
}
