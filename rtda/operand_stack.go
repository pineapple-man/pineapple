package rtda

import (
	"math"
)

// OperandStack 操作数栈的大小是编译器已经确定的
type OperandStack struct {
	// 用于记录栈顶位置
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			size:  0,
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// PushInt 向操作数栈中推入 int 变量
func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].num = val
	o.size++
}

// PopInt 从操作数栈中弹出一个 int 变量
func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].num
}

// PushFloat 向操作数栈中推入 float 变量
func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].num = int32(bits)
	o.size++
}

// PopFloat 从操作数栈中弹出一个 float 变量
func (o *OperandStack) PopFloat() float32 {
	o.size--
	bits := uint32(o.slots[o.size].num)
	return math.Float32frombits(bits)
}

// PushLong 向操作数栈中推入 long 变量
func (o *OperandStack) PushLong(val int64) {
	// save the low 32 bits
	o.slots[o.size].num = int32(val)
	// save the high 32 bits
	o.slots[o.size+1].num = int32(val >> 32)
	o.size += 2
}

// PopLong 从操作数栈中弹出一个 long 变量
func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	low := uint32(o.slots[o.size].num)
	high := uint32(o.slots[o.size+1].num)
	return int64(high)<<32 | int64(low)
}

// PushDouble 向操作数栈中推入 double 变量
func (o *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	o.PushLong(int64(bits))
}

// PopDouble 从操作数栈中弹出一个 double 变量
func (o *OperandStack) PopDouble() float64 {
	bits := uint64(o.PopLong())
	return math.Float64frombits(bits)
}

// PushRef Push a reference to the operand stack
func (o *OperandStack) PushRef(ref *Object) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Object {
	o.size--
	ref := o.slots[o.size].ref
	// 这样做是为了帮助 Go 的垃圾收集器回收 Object 结构体实例
	o.slots[o.size].ref = nil
	return ref
}
