package rtda

import (
	"math"
)

// LocalVars 局部变量表
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make(LocalVars, maxLocals)
	}
	return nil
}

func (l LocalVars) SetInt(index uint, val int32) {
	l[index].num = val
}
func (l LocalVars) GetInt(index uint) int32 {
	return l[index].num
}
func (l LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l[index].num = int32(bits)
}
func (l LocalVars) GetFloat(index uint) float32 {
	bits := uint32(l[index].num)
	return math.Float32frombits(bits)
}
func (l LocalVars) SetLong(index uint, val int64) {
	l[index].num = int32(val)
	l[index+1].num = int32(val >> 32)
}
func (l LocalVars) GetLong(index uint) int64 {
	low := uint32(l[index].num)
	high := uint32(l[index+1].num)
	return int64(high)<<32 | int64(low)
}
func (l LocalVars) SetDouble(index uint, val float64) {
	bits := int64(math.Float64bits(val))
	l.SetLong(index, bits)
}
func (l LocalVars) GetDouble(index uint) float64 {
	bits := l.GetLong(index)
	return math.Float64frombits(uint64(bits))
}
func (l LocalVars) SetRef(index uint, ref *Object) {
	l[index].ref = ref
}
func (l LocalVars) GetRef(index uint) *Object {
	return l[index].ref
}
