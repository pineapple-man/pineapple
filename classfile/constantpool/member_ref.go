package constantpool

import (
	"pineapple/classfile"
)

// ConstantMemberRefInfo 表示 CONSTANT_Fieldref_info、 CONSTANT_Methodref_info、CONSTANT_InterfaceMethodref_info
// 以上三种常量的类型
type ConstantMemberRefInfo struct {
	// 常量池
	constantPool ConstantPool
	// 指向声明方法或字段的类或接口米哦啊舒服，是一个 CONSTANT_NameAndType 的索引项
	classIndex uint16
	// 指向名称以及类型描述符 CONSTANT_NameAndType 的索引项
	nameAndTypeIndex uint16
}

func (c *ConstantMemberRefInfo) readInfo(reader *classfile.ClassReader) {
	c.classIndex = reader.ReadUint16()
	c.nameAndTypeIndex = reader.ReadUint16()
}
func (c *ConstantMemberRefInfo) ClassName() string {
	return c.constantPool.GetClassName(c.classIndex)
}

func (c *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return c.constantPool.getNameAndType(c.nameAndTypeIndex)
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/

// ConstantFieldRefInfo 字段的符号引用
type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

// ConstantMethodRefInfo 类中方法的符号引用
type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

// ConstantInterfaceMethodRefInfo 接口中方法的符号引用
type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}
