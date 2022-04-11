package classfile

import (
	"pineapple/classfile/attribute"
	"pineapple/classfile/constantpool"
)

type MemberInfo struct {
	constantPool    constantpool.ConstantPool // 常量池
	accessFlags     uint16                    // 访问标志
	nameIndex       uint16                    // 字段或方法名
	descriptorIndex uint16                    // 字段或者方法描述符
	attributes      []attribute.Info          // 字段的属性信息
}

// 读取字段表或方法表
func readerMembers(reader *ClassReader, cp constantpool.ConstantPool) []*MemberInfo {
	// 字段或方法表计数器
	memberCount := reader.ReadUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readerMember(reader, cp)
	}
	return members
}
func readerMember(reader *ClassReader, cp constantpool.ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     reader.ReadUint16(),
		nameIndex:       reader.ReadUint16(),
		descriptorIndex: reader.ReadUint16(),
		attributes:      attribute.ReadAttributes(reader, cp),
	}
}

// AccessFlags getter
func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}

// Name 从常量池中查找字段或方法名
func (m *MemberInfo) Name() string {
	return m.constantPool.GetUtf8(m.nameIndex)
}

// Descriptor 从常量池查找字段或方法描述符号
func (m *MemberInfo) Descriptor() string {
	return m.constantPool.GetUtf8(m.descriptorIndex)
}
