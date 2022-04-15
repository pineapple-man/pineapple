package rtda

type Frame struct {
	// 用来实现链表数据结构
	lower *Frame
	// 局部变量表指针
	localVariables LocalVars
	// 操作数栈指针
	operandStack *OperandStack
}

func (f *Frame) LocalVariables() LocalVars {
	return f.localVariables
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

// NewFrame 创建一个新的栈帧
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		lower:          nil,
		localVariables: newLocalVars(maxLocals),
		operandStack:   newOperandStack(maxStack),
	}
}
