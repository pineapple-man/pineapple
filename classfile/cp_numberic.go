package classfile

import (
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/

// ConstantIntegerInfo 与 java 中的int类型常量映射
type ConstantIntegerInfo struct {
	data int32
}

// 先读取一个 uint32 数据，随后将它转型为 int32 类型
func (c ConstantIntegerInfo) readInfo(reader *ClassReader) {
	data := reader.readUint32()
	c.data = int32(data)
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/

// ConstantFloatInfo 使用四个字节存储 IEEE754 但精度浮点数常量
type ConstantFloatInfo struct {
	data float32
}

func (c *ConstantFloatInfo) readInfo(reader *ClassReader) {
	data := reader.readUint32()
	c.data = math.Float32frombits(data)
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantLongInfo struct {
	data int64
}

// 读取一个 unit64 数据，然后将其转型成 int64 类型
func (c *ConstantLongInfo) readInfo(reader *ClassReader) {
	data := reader.readUint64()
	c.data = int64(data)
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/

type ConstantDoubleInfo struct {
	data float64
}

func (c *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	data := reader.readUint64()
	c.data = math.Float64frombits(data)
}
