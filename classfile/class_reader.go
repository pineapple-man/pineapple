package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// ReadUint8 u1
func (c *ClassReader) ReadUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}

// ReadUint16 u2
func (c *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}

// ReadUint32 u4
func (c *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}

// ReadUint64 read uint64 type data
func (c *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}

// ReadUint16s read uint16 table, the size of table size defined by the first uint16 data
func (c *ClassReader) ReadUint16s() []uint16 {
	tableLength := c.ReadUint16()
	s := make([]uint16, tableLength)
	for i := range s {
		s[i] = c.ReadUint16()
	}
	return s
}

// ReadByte read length byte data from the .class file
func (c *ClassReader) ReadByte(length uint32) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}

func (c *ClassReader) ReadBytes(length uint64) []byte {
	bytes := c.data[:length]
	c.data = c.data[length:]
	return bytes
}
