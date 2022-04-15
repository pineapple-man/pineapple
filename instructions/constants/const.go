package constants

import (
	"pineapple/instructions/base"
	"pineapple/rtda"
)

type ACONST_NULL struct {
	base.NoOperandInstructions
}
type DCONST_0 struct {
	base.NoOperandInstructions
}
type DCONST_1 struct {
	base.NoOperandInstructions
}
type FCONST_0 struct {
	base.NoOperandInstructions
}
type FCONST_1 struct {
	base.NoOperandInstructions
}
type FCONST_2 struct {
	base.NoOperandInstructions
}
type ICONST_M1 struct {
	base.NoOperandInstructions
}
type ICONST_0 struct {
	base.NoOperandInstructions
}
type ICONST_1 struct {
	base.NoOperandInstructions
}
type ICONST_2 struct {
	base.NoOperandInstructions
}
type ICONST_3 struct {
	base.NoOperandInstructions
}
type ICONST_4 struct {
	base.NoOperandInstructions
}
type ICONST_5 struct {
	base.NoOperandInstructions
}
type LCONST_0 struct {
	base.NoOperandInstructions
}
type LCONST_1 struct {
	base.NoOperandInstructions
}

// Execute aconst_null 指令把 null 引用推入操作数栈顶
func (A *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Execute dconst_0 指令将 double 类型都 0 推入操作数栈顶
func (D *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}
func (D *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (F *FCONST_0) Execute(frame *rtda.Frame) {
	// TODO implement me
	panic("implement me")
}

func (F *FCONST_1) Execute(frame *rtda.Frame) {
	// TODO implement me
	panic("implement me")
}

func (F *FCONST_2) Execute(frame *rtda.Frame) {
	// TODO implement me
	panic("implement me")
}

// Execute 将 int 型的 -1 推入操作数栈顶
func (I *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (I *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// Execute 将 int 型的 1 推入操作数栈顶
func (I *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

func (I *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

func (I *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

func (I *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

func (I *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

func (L *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(int64(0))
}

func (L *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(int64(1))
}
