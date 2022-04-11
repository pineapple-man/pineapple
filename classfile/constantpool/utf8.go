package constantpool

import (
	"fmt"
	"unicode/utf16"

	"pineapple/classfile"
)

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/

type ConstantUtf8Info struct {
	data string
}

// 读取出 []byte 随后调用 decodeMUTF8() 将其解码成为 Go 字符串
func (c *ConstantUtf8Info) readInfo(reader *classfile.ClassReader) {
	length := reader.ReadUint32()
	bytes := reader.ReadByte(length)
	c.data = decodeMUTF8(bytes)
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytes []byte) string {
	utfLen := len(bytes)
	charArray := make([]uint16, utfLen)
	var c, char2, char3 uint16
	count := 0
	charArrayCount := 0
	for count < utfLen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		charArray[charArrayCount] = c
		charArrayCount++
	}
	for count < utfLen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArray[charArrayCount] = c
			charArrayCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArray[charArrayCount] = c&0x1F<<6 | char2&0x3F
			charArrayCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count-1))
			}
			charArray[charArrayCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrayCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input arount byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	charArray = charArray[0:charArrayCount]
	decodedString := utf16.Decode(charArray)
	return string(decodedString)
}
