package stack

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

type DUP struct {
	base.NoOperandInstructions
}
type DUP_X1 struct {
	base.NoOperandInstructions
}
type DUP_X2 struct {
	base.NoOperandInstructions
}
type DUP2 struct {
	base.NoOperandInstructions
}
type DUP2_X1 struct {
	base.NoOperandInstructions
}
type DUP2_X2 struct {
	base.NoOperandInstructions
}

// Execute dup 指令复制栈顶的单个变量
func (D *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}
