package attribute

import (
	"pineapple/classfile"
	"pineapple/classfile/constantpool"
)

/*
Code_attribute {
	u2 attribute_name_index;
	u4 attribute_length;

	操作数栈的最大深度
	u2 max_stack;

	局部变量表大小
	u2 max_locals;

	u4 code_length;
	u1 code[code_length];
	u2 exception_table_length;
	{
		u2 start_pc;
		u2 end_pc;
		u2 handler_pc;
		u2 catch_type;
	} exception_table[exception_table_length];
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	constantPool   constantpool.ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []Info
}
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *CodeAttribute) readInfo(reader *classfile.ClassReader) {
	c.maxStack = reader.ReadUint16()
	c.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint64()
	c.code = reader.ReadBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.attributes = ReadAttributes(reader, c.constantPool)
}

func readExceptionTable(reader *classfile.ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.ReadUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.ReadUint16(),
			endPc:     reader.ReadUint16(),
			handlerPc: reader.ReadUint16(),
			catchType: reader.ReadUint16(),
		}
	}
	return exceptionTable
}
