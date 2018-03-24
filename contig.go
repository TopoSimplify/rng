package rng

//are two ranges contiguous
func (o *Range) Contiguous(r *Range) bool {
	return (o.I < r.J && o.J == r.I) || (r.I < o.J && r.J == o.I)
}
