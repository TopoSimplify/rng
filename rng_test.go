package rng

import (
	"time"
	"testing"
	"github.com/franela/goblin"
)

func TestRange(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Range", func() {
		g.It("int range", func() {
			g.Timeout(1 * time.Minute)
			rng := NewRange(3, 7)
			xrng := rng.Clone()
			g.Assert(rng.Equals(xrng)).IsTrue()
			g.Assert(rng.I()).Equal(3)
			g.Assert(rng.J()).Equal(7)
			g.Assert(rng.Index(2)).Equal(5)
			g.Assert(rng.Contains(7)).IsTrue()
			g.Assert(rng.Contains(8)).IsFalse()
			g.Assert(rng.Size()).Equal(4)
			g.Assert(rng.Stride()).Equal([]int{3, 4, 5, 6, 7})
			g.Assert(rng.Stride(1)).Equal([]int{3, 4, 5, 6, 7})
			g.Assert(rng.Stride(2)).Equal([]int{3, 5, 7})

			g.Assert(rng.ExclusiveStride()).Equal([]int{4, 5, 6,})
			g.Assert(rng.ExclusiveStride(1)).Equal([]int{4, 5, 6,})
			g.Assert(rng.ExclusiveStride(2)).Equal([]int{4, 6,})

			g.Assert(rng.AsArray()).Equal([2]int{3, 7})
			g.Assert(rng.AsArray()).Equal([2]int{3, 7})

			g.Assert(rng.AsSlice()).Equal([]int{3, 7})
			g.Assert(rng.AsSlice()).Equal([]int{3, 7})
			g.Assert(rng.String()).Equal("Range(i=3, j=7)")
			r := NewRange(0, 9)
			g.Assert(r.Split([]int{3, 5})).Eql([]*Range{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{5, 3, 3, 5})).Eql([]*Range{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{0, 3, 5, 9})).Eql([]*Range{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{0, 3, 5, 9, 13})).Equal([]*Range{{0, 3}, {3, 5}, {5, 9}})
			g.Assert(r.Split([]int{9, 13, 19})).Equal([]*Range{})
		})

		g.It("out of range", func() {
			defer func() {
				r := recover()
				g.Assert(r != nil).IsTrue()
			}()
			g.Timeout(1 * time.Minute)
			rng := NewRange(3, 7)
			xrng := rng.Clone()
			g.Assert(rng.Equals(xrng)).IsTrue()
			g.Assert(rng.Index(0)).Equal(3)
			g.Assert(rng.Index(4)).Equal(7)
			g.Assert(rng.Index(5)).Equal(8)
		})
	})
}
