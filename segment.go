package tracer

// type Segment struct {
// 	sync.RWMutex
// 	Name       string `json:"name,omitempty"`
// 	Type       string `json:"type,omitempty"`
// 	ID         string `json:"id,omitempty"`
// 	TraceID    string `json:"trace_id,omitempty"`
// 	ParentID   string `json:"parent_id,omitempty"`
// 	StartTime  int64  `json:"start_time,omitempty"`
// 	EndTime    int64  `json:"end_time,omitempty"`
// 	InProgress bool   `json:"in_progress,omitempty"`
// }

type Segment interface {
	// sync.RWMutex
	// Name       string `json:"name,omitempty"`
	// Type       string `json:"type,omitempty"`
	// ID         string `json:"id,omitempty"`
	// TraceID    string `json:"trace_id,omitempty"`
	// ParentID   string `json:"parent_id,omitempty"`
	// StartTime  int64  `json:"start_time,omitempty"`
	// EndTime    int64  `json:"end_time,omitempty"`
	// InProgress bool   `json:"in_progress,omitempty"`
	Finish()
}
