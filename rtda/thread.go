package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func (t *Thread) Pc() int {
	return t.pc
}

func (t *Thread) SetPc(pc int) {
	t.pc = pc
}

func (t *Thread) pushFrame(frame *Frame) {
	t.stack.push(frame)
}
func (t *Thread) popFrame() *Frame {
	return t.stack.pop()
}
func (t *Thread) currentFrame() *Frame {
	return t.stack.top()
}

func NewThread() *Thread {
	return &Thread{
		pc:    0,
		stack: newStack(1024),
	}
}
