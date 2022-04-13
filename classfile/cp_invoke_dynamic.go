package classfile

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/

// ConstantMethodTypeInfo 标志方法类型
type ConstantMethodTypeInfo struct {
	// 值必须是对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info 结构，表示方法的描述符
	descriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	c.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/

// ConstantMethodHandleInfo 方法句柄
type ConstantMethodHandleInfo struct {
	// 值必须在 1～9 之间，它决定了方法句柄的类型，方法句柄类型的值表示方法句柄的字节码行为
	referenceKind uint8
	// 值必须是对常量池的有效索引
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	c.referenceKind = reader.readUint8()
	c.referenceIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/

// ConstantInvokeDynamicInfo 表示一个动态方法调用点
type ConstantInvokeDynamicInfo struct {
	// 值必须是对当前 class 文件中引导方法表的 bootstrap_methods[] 数组对有效索引
	bootstrapMethodAttrIndex uint16
	// 值必须是对当前常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_NameAndType_Info 结构，表示方法名和方法描述符
	nameAndTypeIndex uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	c.bootstrapMethodAttrIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}
