// This file was auto-generated by the veyron vdl tool.
// Source: vtrace.vdl

// Package vtrace defines an interface to access Vtrace traces to
// anylize and debug distributed systems.
package vtrace

import (
	"v.io/core/veyron2/services/security/access"

	"v.io/core/veyron2/uniqueid"

	"v.io/core/veyron2/vtrace"

	// The non-user imports are prefixed with "__" to prevent collisions.
	__io "io"
	__veyron2 "v.io/core/veyron2"
	__context "v.io/core/veyron2/context"
	__ipc "v.io/core/veyron2/ipc"
	__vdl "v.io/core/veyron2/vdl"
)

// StoreClientMethods is the client interface
// containing Store methods.
type StoreClientMethods interface {
	// Trace returns the trace that matches the given ID.
	// Will return a NoExists error if no matching trace was found.
	Trace(*__context.T, uniqueid.ID, ...__ipc.CallOpt) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(*__context.T, ...__ipc.CallOpt) (StoreAllTracesCall, error)
}

// StoreClientStub adds universal methods to StoreClientMethods.
type StoreClientStub interface {
	StoreClientMethods
	__ipc.UniversalServiceMethods
}

// StoreClient returns a client stub for Store.
func StoreClient(name string, opts ...__ipc.BindOpt) StoreClientStub {
	var client __ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(__ipc.Client); ok {
			client = clientOpt
		}
	}
	return implStoreClientStub{name, client}
}

type implStoreClientStub struct {
	name   string
	client __ipc.Client
}

func (c implStoreClientStub) c(ctx *__context.T) __ipc.Client {
	if c.client != nil {
		return c.client
	}
	return __veyron2.GetClient(ctx)
}

func (c implStoreClientStub) Trace(ctx *__context.T, i0 uniqueid.ID, opts ...__ipc.CallOpt) (o0 vtrace.TraceRecord, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Trace", []interface{}{i0}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&o0, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c implStoreClientStub) AllTraces(ctx *__context.T, opts ...__ipc.CallOpt) (ocall StoreAllTracesCall, err error) {
	var call __ipc.Call
	if call, err = c.c(ctx).StartCall(ctx, c.name, "AllTraces", nil, opts...); err != nil {
		return
	}
	ocall = &implStoreAllTracesCall{Call: call}
	return
}

// StoreAllTracesClientStream is the client stream for Store.AllTraces.
type StoreAllTracesClientStream interface {
	// RecvStream returns the receiver side of the Store.AllTraces client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() vtrace.TraceRecord
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// StoreAllTracesCall represents the call returned from Store.AllTraces.
type StoreAllTracesCall interface {
	StoreAllTracesClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implStoreAllTracesCall struct {
	__ipc.Call
	valRecv vtrace.TraceRecord
	errRecv error
}

func (c *implStoreAllTracesCall) RecvStream() interface {
	Advance() bool
	Value() vtrace.TraceRecord
	Err() error
} {
	return implStoreAllTracesCallRecv{c}
}

type implStoreAllTracesCallRecv struct {
	c *implStoreAllTracesCall
}

func (c implStoreAllTracesCallRecv) Advance() bool {
	c.c.valRecv = vtrace.TraceRecord{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implStoreAllTracesCallRecv) Value() vtrace.TraceRecord {
	return c.c.valRecv
}
func (c implStoreAllTracesCallRecv) Err() error {
	if c.c.errRecv == __io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implStoreAllTracesCall) Finish() (err error) {
	if ierr := c.Call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

// StoreServerMethods is the interface a server writer
// implements for Store.
type StoreServerMethods interface {
	// Trace returns the trace that matches the given ID.
	// Will return a NoExists error if no matching trace was found.
	Trace(__ipc.ServerContext, uniqueid.ID) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(StoreAllTracesContext) error
}

// StoreServerStubMethods is the server interface containing
// Store methods, as expected by ipc.Server.
// The only difference between this interface and StoreServerMethods
// is the streaming methods.
type StoreServerStubMethods interface {
	// Trace returns the trace that matches the given ID.
	// Will return a NoExists error if no matching trace was found.
	Trace(__ipc.ServerContext, uniqueid.ID) (vtrace.TraceRecord, error)
	// AllTraces returns TraceRecords for all traces the server currently
	// knows about.
	AllTraces(*StoreAllTracesContextStub) error
}

// StoreServerStub adds universal methods to StoreServerStubMethods.
type StoreServerStub interface {
	StoreServerStubMethods
	// Describe the Store interfaces.
	Describe__() []__ipc.InterfaceDesc
}

// StoreServer returns a server stub for Store.
// It converts an implementation of StoreServerMethods into
// an object that may be used by ipc.Server.
func StoreServer(impl StoreServerMethods) StoreServerStub {
	stub := implStoreServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := __ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := __ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implStoreServerStub struct {
	impl StoreServerMethods
	gs   *__ipc.GlobState
}

func (s implStoreServerStub) Trace(ctx __ipc.ServerContext, i0 uniqueid.ID) (vtrace.TraceRecord, error) {
	return s.impl.Trace(ctx, i0)
}

func (s implStoreServerStub) AllTraces(ctx *StoreAllTracesContextStub) error {
	return s.impl.AllTraces(ctx)
}

func (s implStoreServerStub) Globber() *__ipc.GlobState {
	return s.gs
}

func (s implStoreServerStub) Describe__() []__ipc.InterfaceDesc {
	return []__ipc.InterfaceDesc{StoreDesc}
}

// StoreDesc describes the Store interface.
var StoreDesc __ipc.InterfaceDesc = descStore

// descStore hides the desc to keep godoc clean.
var descStore = __ipc.InterfaceDesc{
	Name:    "Store",
	PkgPath: "v.io/core/veyron2/services/mgmt/vtrace",
	Methods: []__ipc.MethodDesc{
		{
			Name: "Trace",
			Doc:  "// Trace returns the trace that matches the given ID.\n// Will return a NoExists error if no matching trace was found.",
			InArgs: []__ipc.ArgDesc{
				{"", ``}, // uniqueid.ID
			},
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // vtrace.TraceRecord
				{"", ``}, // error
			},
			Tags: []__vdl.AnyRep{access.Tag("Debug")},
		},
		{
			Name: "AllTraces",
			Doc:  "// AllTraces returns TraceRecords for all traces the server currently\n// knows about.",
			OutArgs: []__ipc.ArgDesc{
				{"", ``}, // error
			},
			Tags: []__vdl.AnyRep{access.Tag("Debug")},
		},
	},
}

// StoreAllTracesServerStream is the server stream for Store.AllTraces.
type StoreAllTracesServerStream interface {
	// SendStream returns the send side of the Store.AllTraces server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item vtrace.TraceRecord) error
	}
}

// StoreAllTracesContext represents the context passed to Store.AllTraces.
type StoreAllTracesContext interface {
	__ipc.ServerContext
	StoreAllTracesServerStream
}

// StoreAllTracesContextStub is a wrapper that converts ipc.ServerCall into
// a typesafe stub that implements StoreAllTracesContext.
type StoreAllTracesContextStub struct {
	__ipc.ServerCall
}

// Init initializes StoreAllTracesContextStub from ipc.ServerCall.
func (s *StoreAllTracesContextStub) Init(call __ipc.ServerCall) {
	s.ServerCall = call
}

// SendStream returns the send side of the Store.AllTraces server stream.
func (s *StoreAllTracesContextStub) SendStream() interface {
	Send(item vtrace.TraceRecord) error
} {
	return implStoreAllTracesContextSend{s}
}

type implStoreAllTracesContextSend struct {
	s *StoreAllTracesContextStub
}

func (s implStoreAllTracesContextSend) Send(item vtrace.TraceRecord) error {
	return s.s.Send(item)
}
