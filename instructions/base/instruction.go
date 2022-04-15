package base

import (
	"pineapple/rtda"
)

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

// NoOperandInstructions 表示没有操作数的指令
type NoOperandInstructions struct {
}

// FetchOperands 由于没有操作数，所以此方法是一个空方法
func (n *NoOperandInstructions) FetchOperands(reader *ByteCodeReader) {
	// because the no operand instructions have not corresponding operand ,
	//  this method will not execute any statement
}

// BranchInstruction 表示跳转指令
type BranchInstruction struct {
	// 存放跳转偏移量
	offset int
}

func (b *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	// TODO implement me
	panic("implement me")
}

func (b *BranchInstruction) Execute(frame *rtda.Frame) {
	// TODO implement me
	panic("implement me")
}

type Index8Instruction struct {
	Index uint8
}

func (i *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	i.Index = uint8(uint(reader.ReadInt8()))
}

type Index16Instruction struct {
	Index uint16
}

func (i *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	i.Index = uint16(uint(reader.ReadInt16()))
}
