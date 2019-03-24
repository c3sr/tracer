package convert

import (
	"hash/fnv"

	"github.com/Workiva/go-datastructures/augmentedtree"
	model "github.com/uber/jaeger/model/json"
)

type Iterval struct {
	model.Span
}

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func (i Iterval) LowAtDimension(uint64) int64 {
	return int64(i.StartTime)
}

func (i Iterval) HighAtDimension(uint64) int64 {
	return int64(i.StartTime + i.Duration)
}

func (i Iterval) ID() uint64 {
	return hash(string(i.SpanID))
}

func (i Iterval) OverlapsAtDimension(iv augmentedtree.Interval, dimension uint64) bool {
	return i.HighAtDimension(dimension) > iv.LowAtDimension(dimension) &&
		i.LowAtDimension(dimension) < iv.HighAtDimension(dimension)
}

func spanToInterval(s model.Span) augmentedtree.Interval {
	return Iterval{s}
}

func NewSpanTree(spans []model.Span) augmentedtree.Tree {
	tree := augmentedtree.New(1)
	for _, s := range spans {
		tree.Add(spanToInterval(s))
	}
	return tree
}