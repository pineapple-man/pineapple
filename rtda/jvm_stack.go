package rtda

type Stack struct {
	// 保存栈的容量，最多可以容纳多少栈帧
	maxSize uint
	// 保存栈当前大小
	size uint
	// 保存栈顶指针
	_top *Frame
}

// 创建一个栈, maxSize 表示最多可以容纳多少个栈帧
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
		size:    0,
		_top:    nil,
	}
}

// 将一个栈帧压入栈中
func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}
func (s *Stack) pop() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	top := s._top
	s._top = top.lower
	top.lower = nil
	s.size--
	return top
}
func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}
