package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// readUint8 u1
func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}

// readUint16 u2
func (c *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}

// readUint32 u4
func (c *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}

// readUint64 read uint64 type data
func (c *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}

// readUint16s read uint16 table, the size of table size defined by the first uint16 data
func (c *ClassReader) readUint16s() []uint16 {
	tableLength := c.readUint16()
	s := make([]uint16, tableLength)
	for i := range s {
		s[i] = c.readUint16()
	}
	return s
}

// readByte read length byte data from the .class file
func (c *ClassReader) readByte(length uint32) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}

func (c *ClassReader) readBytes(length uint64) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}
