// This file was auto-generated by the veyron idl tool.
// Source: exp.idl

// Package exp is used to test that embedding interfaces works across packages.
// The arith.Calculator idl interface embeds the Exp interface.
package exp

import (
	// The non-user imports are prefixed with "_gen_" to prevent collisions.
	_gen_veyron2 "veyron2"
	_gen_idl "veyron2/idl"
	_gen_ipc "veyron2/ipc"
	_gen_naming "veyron2/naming"
	_gen_rt "veyron2/rt/r"
	_gen_wiretype "veyron2/wiretype"
)

// Exp is the interface the client binds and uses.
// Exp_InternalNoTagGetter is the interface without the TagGetter
// and UnresolveStep methods (both framework-added, rathern than user-defined),
// to enable embedding without method collisions.  Not to be used directly by
// clients.
type Exp_InternalNoTagGetter interface {
	Exp(x float64, opts ..._gen_ipc.ClientCallOpt) (reply float64, err error)
}
type Exp interface {
	_gen_idl.TagGetter
	// UnresolveStep returns the names for the remote service, rooted at the
	// service's immediate namespace ancestor.
	UnresolveStep(opts ..._gen_ipc.ClientCallOpt) ([]string, error)
	Exp_InternalNoTagGetter
}

// ExpService is the interface the server implements.
type ExpService interface {
	Exp(context _gen_ipc.Context, x float64) (reply float64, err error)
}

// BindExp returns the client stub implementing the Exp
// interface.
//
// If no _gen_ipc.Client is specified, the default _gen_ipc.Client in the
// global Runtime is used.
func BindExp(name string, opts ..._gen_ipc.BindOpt) (Exp, error) {
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
			return nil, _gen_idl.ErrUnrecognizedOption
		}
	default:
		return nil, _gen_idl.ErrTooManyOptionsToBind
	}
	stub := &clientStubExp{client: client, name: name}

	return stub, nil
}

// NewServerExp creates a new server stub.
//
// It takes a regular server implementing the ExpService
// interface, and returns a new server stub.
func NewServerExp(server ExpService) interface{} {
	return &ServerStubExp{
		service: server,
	}
}

// clientStubExp implements Exp.
type clientStubExp struct {
	client _gen_ipc.Client
	name   string
}

func (c *clientStubExp) GetMethodTags(method string) []interface{} {
	return GetExpMethodTags(method)
}

func (__gen_c *clientStubExp) Exp(x float64, opts ..._gen_ipc.ClientCallOpt) (reply float64, err error) {
	var call _gen_ipc.ClientCall
	if call, err = __gen_c.client.StartCall(__gen_c.name, "Exp", []interface{}{x}, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

func (c *clientStubExp) UnresolveStep(opts ..._gen_ipc.ClientCallOpt) (reply []string, err error) {
	var call _gen_ipc.ClientCall
	if call, err = c.client.StartCall(c.name, "UnresolveStep", nil, opts...); err != nil {
		return
	}
	if ierr := call.Finish(&reply, &err); ierr != nil {
		err = ierr
	}
	return
}

// ServerStubExp wraps a server that implements
// ExpService and provides an object that satisfies
// the requirements of veyron2/ipc.ReflectInvoker.
type ServerStubExp struct {
	service ExpService
}

func (s *ServerStubExp) GetMethodTags(method string) []interface{} {
	return GetExpMethodTags(method)
}

func (s *ServerStubExp) Signature(call _gen_ipc.ServerCall) (_gen_ipc.ServiceSignature, error) {
	result := _gen_ipc.ServiceSignature{Methods: make(map[string]_gen_ipc.MethodSignature)}
	result.Methods["Exp"] = _gen_ipc.MethodSignature{
		InArgs: []_gen_ipc.MethodArgument{
			{Name: "x", Type: 26},
		},
		OutArgs: []_gen_ipc.MethodArgument{
			{Name: "", Type: 26},
			{Name: "", Type: 65},
		},
	}

	result.TypeDefs = []_gen_idl.AnyData{
		_gen_wiretype.NamedPrimitiveType{Type: 0x1, Name: "error", Tags: []string(nil)}}

	return result, nil
}

func (s *ServerStubExp) UnresolveStep(call _gen_ipc.ServerCall) (reply []string, err error) {
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

func (__gen_s *ServerStubExp) Exp(call _gen_ipc.ServerCall, x float64) (reply float64, err error) {
	reply, err = __gen_s.service.Exp(call, x)
	return
}

func GetExpMethodTags(method string) []interface{} {
	switch method {
	case "Exp":
		return []interface{}{}
	default:
		return nil
	}
}
