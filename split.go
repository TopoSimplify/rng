package rng

import "github.com/intdxdt/iter"

//Split Range at indices
func (o *Range) Split(idxs []int) []*Range {
	idxs = iter.SortedIntsSet(idxs)
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

