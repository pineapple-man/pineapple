package classfile

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/

// ConstantNameAndTypeInfo 字段或方法的符号引用
type ConstantNameAndTypeInfo struct {
	// 指向该字段或方法名称常量的索引
	nameIndex uint16
	// 指向该字段或方法描述符常量的索引
	descriptorIndex uint16
}

func (c ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descriptorIndex = reader.readUint16()
}
