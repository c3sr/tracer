package tracer_test

import (
	"context"
	"testing"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/c3sr/tracer"
	_ "github.com/c3sr/tracer/jaeger"
)

func BenchmarkTracer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		span := tracer.StartSpan(tracer.FULL_TRACE, "test_run")
		span.Finish()
	}
	tracer.Close()
}

func BenchmarkTracerWithContext(b *testing.B) {
	ctx := context.Background()
	for n := 0; n < b.N; n++ {
		var span opentracing.Span
		span, ctx = tracer.StartSpanFromContext(ctx, tracer.FULL_TRACE, "test_run")
		span.Finish()
	}
	tracer.Close()
}
