package classfile

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	// 读取常量信息，需要由具体的常量结构实现。
	readInfo(reader *ClassReader)
}

// 先读出 tag 值，然后调用 newConstantInfo 创建具体的常量
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

// 根据 tag 值创建具体的常量
func newConstantInfo(tag uint8, constantPool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer_info:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float_info:
		return &ConstantFloatInfo{}
	case CONSTANT_Long_info:
		return &ConstantLongInfo{}
	case CONSTANT_Double_info:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8_info:
		return &ConstantUtf8Info{}
	case CONSTANT_String_info:
		return &ConstantStringInfo{constantPool: constantPool}
	case CONSTANT_Class_info:
		return &ConstantClassInfo{constantPool: constantPool}
	case CONSTANT_Fieldref_info:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{constantPool: constantPool}}
	case CONSTANT_Methodref_info:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{constantPool: constantPool}}
	case CONSTANT_Interfacemethodref_info:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{constantPool: constantPool}}
	case CONSTANT_NameAndType_info:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType_info:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle_info:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic_info:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")

	}
}

// ReadConstantPool 读取 class file 中的常量池信息
func ReadConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := int(reader.readUint16())
	constantPool := make([]ConstantInfo, constantPoolCount)
	// 常量池索引从 1 开始
	for i := 1; i < constantPoolCount; i++ {
		constantPool[i] = readConstantInfo(reader, constantPool)
		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			// long 常量和 double 常量 占两个字节
			i++
		}

	}
	return constantPool
}

// 按照索引查找常量
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if constantPoolInfo := c[index]; constantPoolInfo != nil {
		return constantPoolInfo
	}
	panic("Invalid constant pool index !")
}

// 从常量池中查找字段或方法的名字和描述符
func (c ConstantPool) getNameAndType(index uint16) (string, string) {
	nameTypeInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.GetUtf8(nameTypeInfo.nameIndex)
	_type := c.GetUtf8(nameTypeInfo.descriptorIndex)
	return name, _type
}

// GetClassName 从常量池查找类名
func (c ConstantPool) GetClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.GetUtf8(classInfo.nameIndex)
}

// GetUtf8 从常量池中查找 UTF-8 字符串
func (c ConstantPool) GetUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.data
}
