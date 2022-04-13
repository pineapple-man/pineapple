package classfile

import (
	"fmt"
)

type ClassFile struct {
	// magic        uint32 //this field is a flag ,so we need not  save it
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []Info
}

// MinorVersion getter
func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

// MajorVersion getter
func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

// ConstantPool getter
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

// Methods getter
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}

// Fields getter
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

// Parse  : parse []byte data into class file object
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	c.constantPool = ReadConstantPool(reader)
	c.accessFlags = reader.readUint16()
	c.thisClass = reader.readUint16()
	c.superClass = reader.readUint16()
	c.fields = readerMembers(reader, c.constantPool)
	c.methods = readerMembers(reader, c.constantPool)
	c.attributes = ReadAttributes(reader, c.constantPool)
}

// readAndCheckMagic : class 文件的魔数是 0xCAFEBABE，所以读取一个 class 文件的第一步就是检查文件开头的魔数是否正确
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("Error: A JNI error has occurred, please check your installation and " +
			"try again Exception in thread \"main\" java.lang.ClassFormatError: Incompatible" +
			" magic value 1885430635 in class file StringTest")
	}
}

func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 副版本
	c.minorVersion = reader.readUint16()
	// 主版本
	c.majorVersion = reader.readUint16()
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// ClassName : from the constant pool to find the class name
func (c *ClassFile) ClassName() string {
	return c.constantPool.GetClassName(c.thisClass)
}

// SuperClassName : from the constant pool to find this class's super class name
func (c *ClassFile) SuperClassName() string {
	if c.superClass > 0 {
		return c.constantPool.GetClassName(c.superClass)
	}
	// 只有 java.lang.Object 这一个类没有超类
	return ""

}

// InterfaceNames :从常量池中查找接口名
func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))
	for i, cpIndex := range c.interfaces {
		interfaceNames[i] = c.constantPool.GetClassName(cpIndex)
	}
	return interfaceNames
}
