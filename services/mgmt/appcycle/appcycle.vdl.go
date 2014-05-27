// This file was auto-generated by the veyron vdl tool.
// Source: appcycle.vdl

// Package appcycle supports managing the application process.
package appcycle

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt"
	_gen_vdl "veyron2/vdl"
	_gen_wiretype "veyron2/wiretype"
)

// Tick is streamed by Stop to provide the client with a sense of the progress
// of the shutdown.
// The meaning of Progress and Goal are up to the developer (the server provides
// the framework with values for these).  The recommended meanings are:
// - Progress: how far along the shutdown sequence the server is.  This should
//   be a monotonically increasing number.
// - Goal: when Progress reaches this value, the shutdown is expected to
//   complete.  This should not change during a stream, but could change if
//   e.g. new shutdown tasks are triggered that were not forseen at the outset
//   of the shutdown.
type Tick struct {
	Progress int32
	Goal     int32
}

// AppCycle interfaces with the process running a veyron runtime.
// AppCycle is the interface the client binds and uses.
// AppCycle_InternalNoTagGetter is the interface without the TagGetter
// and UnresolveStep methods (both framework-added, rathern than user-defined),
// to enable embedding without method collisions.  Not to be used directly by
// clients.
type AppCycle_InternalNoTagGetter interface {

	// Stop initiates shutdown of the server.  It streams back periodic updates
	// to give the client an idea of how the shutdown is progressing.
	Stop(opts ..._gen_ipc.ClientCallOpt) (reply AppCycleStopStream, err error)
	// ForceStop tells the server to shut down right away.  It can be issued while
	// a Stop is outstanding if for example the client does not want to wait any
	// longer.
	ForceStop(opts ..._gen_ipc.ClientCallOpt) (err error)
}
type AppCycle interface {
	_gen_vdl.TagGetter
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ..._gen_ipc.ClientCallOpt) ([]string, error)
	AppCycle_InternalNoTagGetter
}

// AppCycleService is the interface the server implements.
type AppCycleService interface {

	// Stop initiates shutdown of the server.  It streams back periodic updates
	// to give the client an idea of how the shutdown is progressing.
	Stop(context _gen_ipc.Context, stream AppCycleServiceStopStream) (err error)
	// ForceStop tells the server to shut down right away.  It can be issued while
	// a Stop is outstanding if for example the client does not want to wait any
	// longer.
	ForceStop(context _gen_ipc.Context) (err error)
}

// AppCycleStopStream is the interface for streaming responses of the method
// Stop in the service interface AppCycle.
type AppCycleStopStream interface {

	// Recv returns the next item in the input stream, blocking until
	// an item is available.  Returns io.EOF to indicate graceful end of input.
	Recv() (item Tick, err error)

	// Finish closes the stream and returns the positional return values for
	// call.
	Finish() (err error)

	// Cancel cancels the RPC, notifying the server to stop processing.
	Cancel()
}

// Implementation of the AppCycleStopStream interface that is not exported.
type implAppCycleStopStream struct {
	clientCall _gen_ipc.ClientCall
}

func (c *implAppCycleStopStream) Recv() (item Tick, err error) {
	err = c.clientCall.Recv(&item)
	return
}

func (c *implAppCycleStopStream) Finish() (err error) {
	if ierr := c.clientCall.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *implAppCycleStopStream) Cancel() {
	c.clientCall.Cancel()
}

// AppCycleServiceStopStream is the interface for streaming responses of the method
// Stop in the service interface AppCycle.
type AppCycleServiceStopStream interface {
	// Send places the item onto the output stream, blocking if there is no buffer
	// space available.
	Send(item Tick) error
}

// Implementation of the AppCycleServiceStopStream interface that is not exported.
type implAppCycleServiceStopStream struct {
	serverCall _gen_ipc.ServerCall
}

func (s *implAppCycleServiceStopStream) Send(item Tick) error {
	return s.serverCall.Send(item)
}

// BindAppCycle returns the client stub implementing the AppCycle
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindAppCycle(name string, opts ..._gen_ipc.BindOpt) (AppCycle, error) {
	var client _gen_ipc.Client
	switch len(opts) {
	case 0:
		client = _gen_rt.R().Client()
	case 1:
		switch o := opts[0].(type) {
		case _gen_veyron2.Runtime:
			client = o.Client()
		case _gen_ipc.Client:
			client = o
		default:
			return nil, _gen_vdl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_vdl.ErrTooManyOptionsToBind
	}
	stub := &clientStubAppCycle{client: client, name: name}

	return stub, nil
}

// NewServerAppCycle creates a new server stub.
//
// It takes a regular server implementing the AppCycleService
// interface, and returns a new server stub.
func NewServerAppCycle(server AppCycleService) interface{} {
	return &ServerStubAppCycle{
		service: server,
	}
}

// clientStubAppCycle implements AppCycle.
type clientStubAppCycle struct {
	client _gen_ipc.Client
	name   string
}

func (c *clientStubAppCycle) GetMethodTags(method string) []interface{} {
	return GetAppCycleMethodTags(method)
}

func (__gen_c *clientStubAppCycle) Stop(opts ..._gen_ipc.ClientCallOpt) (reply AppCycleStopStream, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Stop", nil, opts...); err != nil {
		return
	}
	reply = &implAppCycleStopStream{clientCall: call}
	return
}

func (__gen_c *clientStubAppCycle) ForceStop(opts ..._gen_ipc.ClientCallOpt) (err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "ForceStop", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&err); ierr != nil {
		err = ierr
	}
	return
}

func (c *clientStubAppCycle) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = c.client.StartCall(c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubAppCycle wraps a server that implements
// AppCycleService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubAppCycle struct {
	service AppCycleService
}

func (s *ServerStubAppCycle) GetMethodTags(method string) []interface{} {
	return GetAppCycleMethodTags(method)
}

func (s *ServerStubAppCycle) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["ForceStop"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},
	}
	result.Methods["Stop"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 65},
		},

		OutStream: 66,
	}

	result.TypeDefs = []_gen_vdl.Any{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}, _gen_wiretype.StructType{
			[]_gen_wiretype.FieldType{
				_gen_wiretype.FieldType{Type: 0x24, Name: "Progress"},
				_gen_wiretype.FieldType{Type: 0x24, Name: "Goal"},
			},
			"veyron2/services/mgmt/appcycle.Tick", []string(nil)},
	}

	return result, nil
}

func (s *ServerStubAppCycle) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
	if unresolver, ok := s.service.(_gen_ipc.Unresolver); ok {
		return unresolver.UnresolveStep(call)
	}
	if call.Server() == nil {
		return
	}
	var published []string
	if published, err = call.Server().Published(); err != nil || published == nil {
		return
	}
	reply = make([]string, len(published))
	for i, p := range published {
		reply[i] = _gen_naming.Join(p, call.Name())
	}
	return
}

func (__gen_s *ServerStubAppCycle) Stop(call _gen_ipc.ServerCall) (err error) {
	stream := &implAppCycleServiceStopStream{serverCall: call}
	err = __gen_s.service.Stop(call, stream)
	return
}

func (__gen_s *ServerStubAppCycle) ForceStop(call _gen_ipc.ServerCall) (err error) {
	err = __gen_s.service.ForceStop(call)
	return
}

func GetAppCycleMethodTags(method string) []interface{} {
	switch method {
	case "Stop":
		return []interface{}{}
	case "ForceStop":
		return []interface{}{}
	default:
		return nil
	}
}
