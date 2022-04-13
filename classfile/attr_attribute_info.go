package classfile

/*
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
*/

type Info interface {
	readInfo(reader *ClassReader)
}

// ReadAttributes 读取属性表
func ReadAttributes(reader *ClassReader, constantPool ConstantPool) []Info {
	attributesCount := reader.readUint16()
	attributes := make([]Info, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, constantPool)
	}
	return attributes
}

// 读取单个属性
func readAttribute(reader *ClassReader, constantPool ConstantPool) Info {
	attributeNameIndex := reader.readUint16()
	attributeName := constantPool.GetUtf8(attributeNameIndex)
	attributeLength := reader.readUint32()
	attributeInfo := newAttributeInfo(attributeName, attributeLength, constantPool)
	attributeInfo.readInfo(reader)
	return attributeInfo
}

type LocalVariableTableAttribute struct {
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	// TODO implement me
	panic("implement me")
}

func newAttributeInfo(attributeName string, attributeLength uint32, constantPool ConstantPool) Info {
	switch attributeName {
	case "Code":
		return &CodeAttribute{constantPool: constantPool}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{constantPool: constantPool}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attributeName, attributeLength, nil}
	}
}
