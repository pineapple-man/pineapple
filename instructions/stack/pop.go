package stack

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

type POP struct {
	base.NoOperandInstructions
}
type POP2 struct {
	base.NoOperandInstructions
}

// Execute pop 指令将栈顶变量弹出,pop 指令只能用于弹出 int、float 等占用一个操作数栈位置的变量
func (P *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

// Execute pop2 对于 double long 变量在操作数栈中占据两个位置，需要使用 pop2 指令弹出，代码如下
func (P *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
