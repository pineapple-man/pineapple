package classfile

type MemberInfo struct {
	constantPool    ConstantPool // 常量池
	accessFlags     uint16       // 访问标志
	nameIndex       uint16       // 字段或方法名
	descriptorIndex uint16       // 字段或者方法描述符
	attributes      []Info       // 字段的属性信息
}

// 读取字段表或方法表
func readerMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	// 字段或方法表计数器
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readerMember(reader, cp)
	}
	return members
}
func readerMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      ReadAttributes(reader, cp),
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
