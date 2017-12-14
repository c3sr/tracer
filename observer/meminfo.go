package observer

import (
	"runtime"

	"github.com/opentracing-contrib/go-observer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cast"

	olog "github.com/opentracing/opentracing-go/log"
)

var (
	MemInfo otobserver.Observer = memInfo{}
)

type memInfo struct{}

// OnStartSpan creates a new memInfo for the span
func (o memInfo) OnStartSpan(sp opentracing.Span, operationName string, options opentracing.StartSpanOptions) (otobserver.SpanObserver, bool) {
	if operationName != "Predict" {
		return nil, false
	}
	return newMemInfoSpan(sp, options)
}

// SpanDummy collects perfevent metrics
type memInfoSpan struct {
	sp opentracing.Span
}

// NewSpanDummy creates a new SpanDummy that can emit perfevent
// metrics
func newMemInfoSpan(s opentracing.Span, opts opentracing.StartSpanOptions) (*memInfoSpan, bool) {
	so := &memInfoSpan{
		sp: s,
	}
	memStats := getCPUMemUsage()
	v, _ := mem.VirtualMemory()
	s.LogFields(
		olog.String("start_mem_alloc", cast.ToString(memStats.Alloc)),
		olog.String("start_mem_total_alloc", cast.ToString(memStats.TotalAlloc)),
		olog.String("start_mem_heap_alloc", cast.ToString(memStats.HeapAlloc)),
		olog.String("start_mem_heap_sys", cast.ToString(memStats.HeapSys)),
	)
	s.LogFields(
		olog.String("start_mem_sys_available", cast.ToString(v.Available)),
		olog.String("start_mem_sys_free", cast.ToString(v.Free)),
		olog.String("start_mem_sys_total", cast.ToString(v.Total)),
	)

	return so, true
}

func (so *memInfoSpan) OnSetOperationName(operationName string) {
}

func (so *memInfoSpan) OnSetTag(key string, value interface{}) {
}

func (so *memInfoSpan) OnFinish(options opentracing.FinishOptions) {
	memStats := getCPUMemUsage()
	v, _ := mem.VirtualMemory()
	so.sp.LogFields(
		olog.String("finish_mem_alloc", cast.ToString(memStats.Alloc)),
		olog.String("finish_mem_total_alloc", cast.ToString(memStats.TotalAlloc)),
		olog.String("finish_mem_heap_alloc", cast.ToString(memStats.HeapAlloc)),
		olog.String("finish_mem_heap_sys", cast.ToString(memStats.HeapSys)),
	)
	so.sp.LogFields(
		olog.String("finish_mem_sys_available", cast.ToString(v.Available)),
		olog.String("finish_mem_sys_free", cast.ToString(v.Free)),
		olog.String("finish_mem_sys_total", cast.ToString(v.Total)),
	)
}

func getCPUMemUsage() runtime.MemStats {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return mem
}
