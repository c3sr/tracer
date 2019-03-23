package main

import (
	"C"
)

import (
	"context"
	"sync"
	"unsafe"

	"github.com/opentracing/opentracing-go"
	"github.com/rai-project/tracer"
	_ "github.com/rai-project/tracer/jaeger"
	// _ "github.com/rai-project/tracer/noop"
	// _ "github.com/rai-project/tracer/zipkin"
)

var spans sync.Map

//go:nosplit
func fromSpan(sp opentracing.Span) uintptr {
	return (uintptr)(unsafe.Pointer(&sp))
}

//go:nosplit
func toSpan(sp uintptr) opentracing.Span {
	return *((*opentracing.Span)(unsafe.Pointer(sp)))
}

//go:nosplit
func fromContext(ctx context.Context) uintptr {
	return (uintptr)(unsafe.Pointer(&ctx))
}

//go:nosplit
func toContext(ctx uintptr) context.Context {
	return *((*context.Context)(unsafe.Pointer(ctx)))
}

//export SpanStart
func SpanStart(lvl int32, operationName string) uintptr {
	sp := tracer.StartSpan(tracer.Level(lvl), operationName)
	spanPtr := fromSpan(sp)
	spans.Store(spanPtr, sp)
	return spanPtr
}

//export SpanAddTag
func SpanAddTag(spPtr uintptr, key, val string) {
	sp := toSpan(spPtr)
	sp.SetTag(key, val)
}

//export SpanAddTags
func SpanAddTags(spPtr uintptr, len int, keys []string, vals []string) {
	sp := toSpan(spPtr)
	for ii := 0; ii < len; ii++ {
		sp.SetTag(keys[ii], vals[ii])
	}
}

//export SpanAddArgumentsTag
func SpanAddArgumentsTag(spPtr uintptr, len int, keys []string, vals []string) {
	sp := toSpan(spPtr)
	args := make(map[string]string, len)
	for ii := 0; ii < len; ii++ {
		args[keys[ii]] = vals[ii]
	}
	sp.SetTag("arguments", args)
}

//export SpanFinish
func SpanFinish(spPtr uintptr) {
	var sp opentracing.Span
	if e, ok := spans.Load(spPtr); ok {
		sp = e.(opentracing.Span)
	}
	if sp == nil {
		return
	}
	sp.Finish()
	spans.Delete(spPtr)
}

//export StartSpanFromContext
func StartSpanFromContext(inCtx uintptr, lvl int32, operationName string, tags map[string]string) (uintptr, uintptr) {
	sp, ctx := tracer.StartSpanFromContext(toContext(inCtx), tracer.Level(lvl), operationName, cTags(tags))
	return fromSpan(sp), fromContext(ctx)
}
