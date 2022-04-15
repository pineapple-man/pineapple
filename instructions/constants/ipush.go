package constants

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

// BIPUSH bipush 指令从操作数中获取一个 byte 型整数，扩展成 int 型，然后推入栈顶
type BIPUSH struct {
	val int8
}

// sipush 指令从操作数中获取一个 short 型整数，扩展成 int 型，然后推入栈顶
type SIPUSH struct {
	val int16
}

func (s *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	s.val = reader.ReadInt16()
}

func (s *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}

func (b *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	b.val = reader.ReadInt8()
}

func (b *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(b.val)
	frame.OperandStack().PushInt(i)
}
