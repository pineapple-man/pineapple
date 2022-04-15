package base

type ByteCodeReader struct {
	// 存放字节码
	code []byte
	// 记录读取到了哪个字节
	pc int
}

func (b *ByteCodeReader) Reset(code []byte, pc int) {
	b.code = code
	b.pc = pc
}

func (b *ByteCodeReader) ReadUint8() uint8 {
	i := b.code[b.pc]
	b.pc++
	return i
}
func (b *ByteCodeReader) ReadUint16() uint16 {
	byte1 := uint16(b.ReadUint8())
	byte2 := uint16(b.ReadUint8())
	return byte1<<8 | byte2
}
func (b *ByteCodeReader) ReadUint32() uint32 {
	byte1 := uint32(b.ReadUint8())
	byte2 := uint32(b.ReadUint8())
	byte3 := uint32(b.ReadUint8())
	byte4 := uint32(b.ReadUint8())
	return byte1<<24 | byte2<<16 | byte3<<8 | byte4
}

func (b *ByteCodeReader) ReadInt8() int8 {
	return int8(b.ReadUint8())
}

func (b *ByteCodeReader) ReadInt16() int16 {
	return int16(b.ReadUint16())
}
func (b *ByteCodeReader) ReadInt32() int32 {
	return int32(b.ReadUint32())
}
