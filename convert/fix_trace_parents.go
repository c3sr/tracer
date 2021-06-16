package convert

import (
	"github.com/pkg/errors"
	model "github.com/jaegertracing/jaeger/model/json"
)

func FixTrace(trace model.Trace) (model.Trace, error) {
	return FixParentRelationship(trace)
}

func FixParentRelationship(trace model.Trace) (model.Trace, error) {
	tree, err := NewIntervalTree(trace)
	if err != nil {
		return model.Trace{}, errors.Wrap(err, "failed to create interval tree for trace")
	}
	return tree.FixParentRelationship()
}
