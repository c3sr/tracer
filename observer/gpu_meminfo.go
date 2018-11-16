// +build !nogpu
// +build linux
// +build !arm64
// +build cgo

package observer

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/opentracing-contrib/go-observer"
	opentracing "github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
	"github.com/rai-project/config"
	nvml "github.com/rai-project/nvml-go"
	"github.com/spf13/cast"
)

var (
	GPUMemInfo otobserver.Observer
)

func init() {
	config.BeforeInit(func() {
		GPUMemInfo = newGPUMemInfo()
	})
}

type gpuMemInfo struct {
	count   int
	handles []nvml.DeviceHandle
}

func newGPUMemInfo() *gpuMemInfo {
	err := nvml.Init()
	if err != nil {
		panic(pp.Sprint("failed to init nvml = ", err))
	}

	count, err := nvml.DeviceCount()
	if err != nil {
		panic(err)
	}
	handles := make([]nvml.DeviceHandle, count)
	for ii := range handles {
		handle, err := nvml.DeviceGetHandleByIndex(ii)
		if err != nil {
			panic(pp.Sprint("failed to create device handle = ", err))
		}
		handles[ii] = handle
	}

	return &gpuMemInfo{
		count:   count,
		handles: handles,
	}
}

// OnStartSpan creates a new gpuMemInfo for the span
func (o *gpuMemInfo) OnStartSpan(sp opentracing.Span, operationName string, options opentracing.StartSpanOptions) (otobserver.SpanObserver, bool) {
	if operationName != "Predict" {
		return nil, false
	}
	return newGPUMemInfoSpan(o, sp, options)
}

// SpanDummy collects perfevent metrics
type gpuMemInfoSpan struct {
	*gpuMemInfo
	sp opentracing.Span
}

// NewSpanDummy creates a new SpanDummy that can emit perfevent
// metrics
func newGPUMemInfoSpan(info *gpuMemInfo, s opentracing.Span, opts opentracing.StartSpanOptions) (*gpuMemInfoSpan, bool) {

	so := &gpuMemInfoSpan{
		gpuMemInfo: info,
		sp:         s,
	}
	for ii, handle := range so.handles {
		meminfo, err := nvml.DeviceMemoryInformation(handle)
		if err != nil {
			continue
		}
		prefix := fmt.Sprintf("start_gpu[%v]_", ii)
		s.LogFields(
			olog.String(prefix+"mem_used", cast.ToString(meminfo.Used)),
			olog.String(prefix+"mem_free", cast.ToString(meminfo.Free)),
			olog.String(prefix+"mem_total", cast.ToString(meminfo.Total)),
		)
	}

	return so, true
}

func (so *gpuMemInfoSpan) OnSetOperationName(operationName string) {
}

func (so *gpuMemInfoSpan) OnSetTag(key string, value interface{}) {
}

func (so *gpuMemInfoSpan) OnFinish(options opentracing.FinishOptions) {
	for ii, handle := range so.handles {
		meminfo, err := nvml.DeviceMemoryInformation(handle)
		if err != nil {
			continue
		}
		prefix := fmt.Sprintf("finish_gpu[%v]_", ii)
		so.sp.LogFields(
			olog.String(prefix+"mem_used", cast.ToString(meminfo.Used)),
			olog.String(prefix+"mem_free", cast.ToString(meminfo.Free)),
			olog.String(prefix+"mem_total", cast.ToString(meminfo.Total)),
		)
	}
}
