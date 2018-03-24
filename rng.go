package rng

import (
	"fmt"
)

//Range
type Range struct {
	I int
	J int
}

//New Range
func NewRange(i, j int) *Range {
	return &Range{i, j}
}

//Stringer interface
func (o *Range) String() string {
	return fmt.Sprintf("Range(i=%v, j=%v)", o.I, o.J)
}


//clone Range
func (o *Range) Clone() *Range {
	return NewRange(o.I, o.J)
}

//compare equality of two ranges
func (o *Range) Equals(r *Range) bool {
	return (o.I == r.I) && (o.J == r.J)
}


//Size
func (o *Range) Size() int {
	return o.J - o.I
}
