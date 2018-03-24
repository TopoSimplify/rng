package rng

//As Array
func (o *Range) AsArray() [2]int {
	return [2]int{o.I, o.J}
}

//As Slice
func (o *Range) AsSlice() []int {
	ar := o.AsArray()
	return ar[:]
}

