package rng

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Rng", func() {
		g.It("int range", func() {
			g.Timeout(1 * time.Minute)
			var arng = Range(7, 12)
			var brng = Range(0, 3)
			var crng = Range(2, 5)
			var rng = Range(3, 7)
			var xrng = rng
			g.Assert(arng.Contiguous(xrng)).IsTrue()
			g.Assert(brng.Contiguous(rng)).IsTrue()
			g.Assert(crng.Contiguous(rng)).IsFalse()
			g.Assert(rng.Contiguous(xrng)).IsFalse()

			g.Assert(rng.Equals(xrng)).IsTrue()
			g.Assert(rng.I).Equal(3)
			g.Assert(rng.J).Equal(7)

			g.Assert(rng.Size()).Equal(4)
			g.Assert(rng.Stride()).Equal([]int{3, 4, 5, 6, 7})
			g.Assert(rng.Stride(1)).Equal([]int{3, 4, 5, 6, 7})
			g.Assert(rng.Stride(2)).Equal([]int{3, 5, 7})

			g.Assert(rng.ExclusiveStride()).Equal([]int{4, 5, 6,})
			g.Assert(rng.ExclusiveStride(1)).Equal([]int{4, 5, 6,})
			g.Assert(rng.ExclusiveStride(2)).Equal([]int{4, 6,})

			g.Assert(rng.AsArray()).Equal([2]int{3, 7})

			g.Assert(rng.AsSlice()).Equal([]int{3, 7})
			g.Assert(rng.String()).Equal("Rng(i=3, j=7)")
			r := Range(0, 9)
			g.Assert(r.Split([]int{3, 5})).Eql([]Rng{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{5, 3, 3, 5})).Eql([]Rng{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{0, 3, 5, 9})).Eql([]Rng{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{0, 3, 5, 9, 13})).Equal([]Rng{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{9, 13, 19})).Equal([]Rng{})
		})
	})
}
