package stack

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

// SWAP the top two operand stack values
type SWAP struct {
	base.NoOperandInstructions
}

// Execute 交换栈顶元素
func (S *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	top := stack.PopSlot()
	subTop := stack.PopSlot()
	stack.PushSlot(top)
	stack.PushSlot(subTop)
}
