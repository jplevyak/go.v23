// This file was auto-generated by the veyron vdl tool.
// Source: appcycle.vdl

// Package appcycle supports managing the application process.
package appcycle

import (
	// VDL system imports
	"io"
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/ipc"
	"v.io/v23/vdl"
)

// Task is streamed by Stop to provide the client with a sense of the progress
// of the shutdown.
// The meaning of Progress and Goal are up to the developer (the server provides
// the framework with values for these).  The recommended meanings are:
// - Progress: how far along the shutdown sequence the server is.  This should
//   be a monotonically increasing number.
// - Goal: when Progress reaches this value, the shutdown is expected to
//   complete.  This should not change during a stream, but could change if
//   e.g. new shutdown tasks are triggered that were not forseen at the outset
//   of the shutdown.
type Task struct {
	Progress int32
	Goal     int32
}

func (Task) __VDLReflect(struct {
	Name string "v.io/v23/services/mgmt/appcycle.Task"
}) {
}

func init() {
	vdl.Register((*Task)(nil))
}

// AppCycleClientMethods is the client interface
// containing AppCycle methods.
//
// AppCycle interfaces with the process running a veyron runtime.
type AppCycleClientMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(*context.T, ...ipc.CallOpt) (AppCycleStopClientCall, error)
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(*context.T, ...ipc.CallOpt) error
}

// AppCycleClientStub adds universal methods to AppCycleClientMethods.
type AppCycleClientStub interface {
	AppCycleClientMethods
	ipc.UniversalServiceMethods
}

// AppCycleClient returns a client stub for AppCycle.
func AppCycleClient(name string, opts ...ipc.BindOpt) AppCycleClientStub {
	var client ipc.Client
	for _, opt := range opts {
		if clientOpt, ok := opt.(ipc.Client); ok {
			client = clientOpt
		}
	}
	return implAppCycleClientStub{name, client}
}

type implAppCycleClientStub struct {
	name   string
	client ipc.Client
}

func (c implAppCycleClientStub) c(ctx *context.T) ipc.Client {
	if c.client != nil {
		return c.client
	}
	return v23.GetClient(ctx)
}

func (c implAppCycleClientStub) Stop(ctx *context.T, opts ...ipc.CallOpt) (ocall AppCycleStopClientCall, err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "Stop", nil, opts...); err != nil {
		return
	}
	ocall = &implAppCycleStopClientCall{ClientCall: call}
	return
}

func (c implAppCycleClientStub) ForceStop(ctx *context.T, opts ...ipc.CallOpt) (err error) {
	var call ipc.ClientCall
	if call, err = c.c(ctx).StartCall(ctx, c.name, "ForceStop", nil, opts...); err != nil {
		return
	}
	err = call.Finish()
	return
}

// AppCycleStopClientStream is the client stream for AppCycle.Stop.
type AppCycleStopClientStream interface {
	// RecvStream returns the receiver side of the AppCycle.Stop client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() Task
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// AppCycleStopClientCall represents the call returned from AppCycle.Stop.
type AppCycleStopClientCall interface {
	AppCycleStopClientStream
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

type implAppCycleStopClientCall struct {
	ipc.ClientCall
	valRecv Task
	errRecv error
}

func (c *implAppCycleStopClientCall) RecvStream() interface {
	Advance() bool
	Value() Task
	Err() error
} {
	return implAppCycleStopClientCallRecv{c}
}

type implAppCycleStopClientCallRecv struct {
	c *implAppCycleStopClientCall
}

func (c implAppCycleStopClientCallRecv) Advance() bool {
	c.c.valRecv = Task{}
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implAppCycleStopClientCallRecv) Value() Task {
	return c.c.valRecv
}
func (c implAppCycleStopClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implAppCycleStopClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// AppCycleServerMethods is the interface a server writer
// implements for AppCycle.
//
// AppCycle interfaces with the process running a veyron runtime.
type AppCycleServerMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(AppCycleStopServerCall) error
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(ipc.ServerCall) error
}

// AppCycleServerStubMethods is the server interface containing
// AppCycle methods, as expected by ipc.Server.
// The only difference between this interface and AppCycleServerMethods
// is the streaming methods.
type AppCycleServerStubMethods interface {
	// Stop initiates shutdown of the server.  It streams back periodic
	// updates to give the client an idea of how the shutdown is
	// progressing.
	Stop(*AppCycleStopServerCallStub) error
	// ForceStop tells the server to shut down right away.  It can be issued
	// while a Stop is outstanding if for example the client does not want
	// to wait any longer.
	ForceStop(ipc.ServerCall) error
}

// AppCycleServerStub adds universal methods to AppCycleServerStubMethods.
type AppCycleServerStub interface {
	AppCycleServerStubMethods
	// Describe the AppCycle interfaces.
	Describe__() []ipc.InterfaceDesc
}

// AppCycleServer returns a server stub for AppCycle.
// It converts an implementation of AppCycleServerMethods into
// an object that may be used by ipc.Server.
func AppCycleServer(impl AppCycleServerMethods) AppCycleServerStub {
	stub := implAppCycleServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := ipc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := ipc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implAppCycleServerStub struct {
	impl AppCycleServerMethods
	gs   *ipc.GlobState
}

func (s implAppCycleServerStub) Stop(call *AppCycleStopServerCallStub) error {
	return s.impl.Stop(call)
}

func (s implAppCycleServerStub) ForceStop(call ipc.ServerCall) error {
	return s.impl.ForceStop(call)
}

func (s implAppCycleServerStub) Globber() *ipc.GlobState {
	return s.gs
}

func (s implAppCycleServerStub) Describe__() []ipc.InterfaceDesc {
	return []ipc.InterfaceDesc{AppCycleDesc}
}

// AppCycleDesc describes the AppCycle interface.
var AppCycleDesc ipc.InterfaceDesc = descAppCycle

// descAppCycle hides the desc to keep godoc clean.
var descAppCycle = ipc.InterfaceDesc{
	Name:    "AppCycle",
	PkgPath: "v.io/v23/services/mgmt/appcycle",
	Doc:     "// AppCycle interfaces with the process running a veyron runtime.",
	Methods: []ipc.MethodDesc{
		{
			Name: "Stop",
			Doc:  "// Stop initiates shutdown of the server.  It streams back periodic\n// updates to give the client an idea of how the shutdown is\n// progressing.",
		},
		{
			Name: "ForceStop",
			Doc:  "// ForceStop tells the server to shut down right away.  It can be issued\n// while a Stop is outstanding if for example the client does not want\n// to wait any longer.",
		},
	},
}

// AppCycleStopServerStream is the server stream for AppCycle.Stop.
type AppCycleStopServerStream interface {
	// SendStream returns the send side of the AppCycle.Stop server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item Task) error
	}
}

// AppCycleStopServerCall represents the context passed to AppCycle.Stop.
type AppCycleStopServerCall interface {
	ipc.ServerCall
	AppCycleStopServerStream
}

// AppCycleStopServerCallStub is a wrapper that converts ipc.StreamServerCall into
// a typesafe stub that implements AppCycleStopServerCall.
type AppCycleStopServerCallStub struct {
	ipc.StreamServerCall
}

// Init initializes AppCycleStopServerCallStub from ipc.StreamServerCall.
func (s *AppCycleStopServerCallStub) Init(call ipc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the AppCycle.Stop server stream.
func (s *AppCycleStopServerCallStub) SendStream() interface {
	Send(item Task) error
} {
	return implAppCycleStopServerCallSend{s}
}

type implAppCycleStopServerCallSend struct {
	s *AppCycleStopServerCallStub
}

func (s implAppCycleStopServerCallSend) Send(item Task) error {
	return s.s.Send(item)
}
