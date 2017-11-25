package rng

import (
	"fmt"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/iter"
	"github.com/intdxdt/sset"
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

//Get Index from I
func (o *Range) Index(index int) int {
	var k = o.I + index
	if k > o.J {
		panic("index out of bounds, i <= k <= j")
	}
	return k
}

//clone Range
func (o *Range) Clone() *Range {
	return NewRange(o.I, o.J)
}

//compare equality of two ranges
func (o *Range) Equals(r *Range) bool {
	return (o.I == r.I) && (o.J == r.J)
}

//compare equality of two ranges
func (o *Range) Contiguous(r *Range) bool {
	return (o.I < r.J && o.J == r.I) || (r.I < o.J && r.J == o.I)
}

//as segment
func (o *Range) Contains(idx int) bool {
	return idx == o.I || idx == o.J
}

//As Array
func (o *Range) AsArray() [2]int {
	return [2]int{o.I, o.J}
}

//As Slice
func (o *Range) AsSlice() []int {
	ar := o.AsArray()
	return ar[:]
}

//Size
func (o *Range) Size() int {
	return o.J - o.I
}

//Stride
func (o *Range) Stride(step ...int) []int {
	var s = 1
	if len(step) > 0 {
		s = step[0]
	}
	return iter.NewGenerator(o.I, o.J+1, s).Values()
}

//Exclusive stride
func (o *Range) ExclusiveStride(step ...int) []int {
	var s = 1
	if len(step) > 0 {
		s = step[0]
	}
	return iter.NewGenerator(o.I+1, o.J, s).Values()
}

//Split Range at indices
func (o *Range) Split(idxs []int) []*Range {
	var idxset = sset.NewSSet(cmp.Int)
	for _, v := range idxs {
		idxset.Add(v)
	}
	idxs = make([]int, 0, idxset.Size())
	for _, o := range idxset.Values() {
		idxs = append(idxs, o.(int))
	}

	var i, j = o.I, o.J
	var sub = make([]*Range, 0)
	for _, idx := range idxs {
		if i < idx && idx < j {
			s := i
			if len(sub) > 0 {
				s = sub[len(sub)-1].J
			}
			sub = append(sub, NewRange(s, idx))
		}
	}
	//close rest of split
	if len(sub) > 0 {
		e := sub[len(sub)-1].J
		if e < j {
			sub = append(sub, NewRange(e, j))
		}
	}
	return sub
}
