package rng

//Get Index from I
func (o Rng) Index(index int) int {
	var k = o.I + index
	if k > o.J {
		panic("index out of bounds, i <= k <= j")
	}
	return k
}
