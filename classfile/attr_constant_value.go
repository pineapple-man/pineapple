package classfile

/*
ConstantValue 是定长属性，只会出现在field_info结构中，用于表示常量表达式的值
attribute_length 的值必须是2
ConstantValue_attribute {
	u2 attribute_name_index;
	u4 attribute_length;
	u2 constantvalue_index;
}
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (c *ConstantValueAttribute) readInfo(reader *ClassReader) {
	c.constantValueIndex = reader.readUint16()
}
func (c *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return c.constantValueIndex
}
