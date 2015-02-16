// This file was auto-generated by the veyron vdl tool.
// Source: types.vdl

package vtrace

import (
	// VDL system imports
	"v.io/core/veyron2/vdl"

	// VDL user imports
	"time"
	"v.io/core/veyron2/uniqueid"
	_ "v.io/core/veyron2/vdl/vdlroot/src/time"
)

type TraceRecord struct {
	ID    uniqueid.Id
	Spans []SpanRecord
}

func (TraceRecord) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.TraceRecord"
}) {
}

// An Annotation represents data that is relevant at a specific moment.
// They can be attached to spans to add useful debugging information.
type Annotation struct {
	// When the annotation was added.
	When time.Time
	// The annotation message.
	// TODO(mattr): Allow richer annotations.
	Message string
}

func (Annotation) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.Annotation"
}) {
}

// A SpanRecord is the wire format for a Span.
type SpanRecord struct {
	ID     uniqueid.Id // The ID of the Span.
	Parent uniqueid.Id // The ID of this Span's parent.
	Name   string      // The Name of this span.
	Start  time.Time   // The start time of this span.
	End    time.Time   // The end time of this span.
	// A series of annotations.
	Annotations []Annotation
}

func (SpanRecord) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.SpanRecord"
}) {
}

type TraceFlags int32

func (TraceFlags) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.TraceFlags"
}) {
}

// Request is the object that carries trace informtion between processes.
type Request struct {
	SpanID  uniqueid.Id // The ID of the span that originated the RPC call.
	TraceID uniqueid.Id // The ID of the trace this call is a part of.
	Flags   TraceFlags
}

func (Request) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.Request"
}) {
}

type Response struct {
	// Flags give options for trace collection, the client should alter its
	// collection for this trace according to the flags sent back from the
	// server.
	Flags TraceFlags
	// Trace is collected trace data.  This may be empty.
	Trace TraceRecord
}

func (Response) __VDLReflect(struct {
	Name string "v.io/core/veyron2/vtrace.Response"
}) {
}

func init() {
	vdl.Register((*TraceRecord)(nil))
	vdl.Register((*Annotation)(nil))
	vdl.Register((*SpanRecord)(nil))
	vdl.Register((*TraceFlags)(nil))
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))
}

const Empty = TraceFlags(0)

const CollectInMemory = TraceFlags(1)
