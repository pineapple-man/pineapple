package constants

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

// NOP 代表空指令(nop instruction)
type NOP struct {
	base.NoOperandInstructions
}

func (N *NOP) Execute(frame *rtda.Frame) {
	// nop instruction will not execute any statement
}
